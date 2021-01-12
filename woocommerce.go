package woocommerce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	UserAgent            = "woocommerce/1.0.0"
	defaultHttpTimeout   = 10
	defaultApiPathPrefix = "/wp-json/wc/v3"
	defaultVersion       = "v3"
)

var (
	apiVersionRegex = regexp.MustCompile(`^v[0-9]{2}$`)
)

type App struct {
	CustomerKey    string
	CustomerSecret string
	AppName        string
	UserId         string
	Scope          string
	ReturnUrl      string
	CallbackUrl    string
	Client         *Client
}

type RateLimitInfo struct {
	RequestCount      int
	BucketSize        int
	RetryAfterSeconds float64
}

type Client struct {
	Client     *http.Client
	app        App
	version    string
	log        LeveledLoggerInterface
	baseURL    *url.URL
	pathPrefix string
	token      string

	// max number of retries, defaults to 0 for no retries see WithRetry option
	retries  int
	attempts int

	RateLimits RateLimitInfo
	Product    ProductService
	Order      OrderService
	OrderNote  OrderNoteService
}

// NewClient returns a new WooCommerce API client with an already authenticated shopname and
// token. The shopName parameter is the shop's wooCommerce website domain,
// e.g. "shop.gitvim.com"
// a.NewClient(shopName, token, opts) is equivalent to NewClient(a, shopName, token, opts)
func (a App) NewClient(shopName string, opts ...Option) *Client {
	return NewClient(a, shopName, opts...)
}

// NewClient Returns a new WooCommerce API client with an already authenticated shopname and
// token. The shopName parameter is the shop's wooCommerce website domain,
// e.g. "shop.gitvim.com"
func NewClient(app App, shopName string, opts ...Option) *Client {
	baseURL, err := url.Parse(ShopBaseURL(shopName))
	if err != nil {
		panic(err)
	}
	c := &Client{
		Client: &http.Client{
			Timeout: time.Second * defaultHttpTimeout,
		},
		log:        &LeveledLogger{},
		app:        app,
		baseURL:    baseURL,
		version:    defaultVersion,
		pathPrefix: defaultApiPathPrefix,
	}
	c.Product = &ProductServiceOp{client: c}
	c.Order = &OrderServiceOp{client: c}
	c.OrderNote = &OrderNoteServiceOp{client: c}
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// ShopBaseURL return a shop's base https base url
func ShopBaseURL(shopName string) string {
	return fmt.Sprintf("https://%s", shopName)
}

// Do sends an API request and populates the given interface with the parsed
// response. It does not make much sense to call Do without a prepared
// interface instance.
func (c *Client) Do(req *http.Request, v interface{}) error {
	_, err := c.doGetHeaders(req, v)
	if err != nil {
		return err
	}

	return nil
}

// doGetHeaders executes a request, decoding the response into `v` and also returns any response headers.
func (c *Client) doGetHeaders(req *http.Request, v interface{}) (http.Header, error) {
	var resp *http.Response
	var err error
	retries := c.retries
	c.attempts = 0
	c.logRequest(req)

	for {
		c.attempts++
		resp, err = c.Client.Do(req)

		c.logResponse(resp)
		if err != nil {
			return nil, err //http client errors, not api responses
		}

		respErr := CheckResponseError(resp)
		if respErr == nil {
			println("break", respErr)
			break // no errors, break out of the retry loop
		}

		// retry scenario, close resp and any continue will retry
		resp.Body.Close()

		if retries <= 1 {
			return nil, respErr
		}

		if rateLimitErr, isRetryErr := respErr.(RateLimitError); isRetryErr {
			wait := time.Duration(rateLimitErr.RetryAfter) * time.Second
			c.log.Debugf("rate limited waiting %s", wait.String())
			time.Sleep(wait)
			retries--
			continue
		}

		var doRetry bool
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			c.log.Debugf("service unavailable, retrying")
			doRetry = true
			retries--
		}

		if doRetry {
			continue
		}

		// no retry attempts, just return the err
		return nil, respErr
	}

	c.logResponse(resp)
	defer resp.Body.Close()

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&v)
		if err != nil {
			return nil, err
		}
	}

	return resp.Header, nil
}

// ResponseDecodingError occurs when the response body from WooCommerce could
// not be parsed.
type ResponseDecodingError struct {
	Body    []byte
	Message string
	Status  int
}

func (e ResponseDecodingError) Error() string {
	return e.Message
}

func CheckResponseError(r *http.Response) error {
	if http.StatusOK <= r.StatusCode && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}

	// Create an anonoymous struct to parse the JSON data into.
	woocommerceError := struct {
		Error  string      `json:"error"`
		Errors interface{} `json:"errors"`
	}{}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// empty body, this probably means WooCommerce returned an error with no body
	// we'll handle that error in wrapSpecificError()
	if len(bodyBytes) > 0 {
		err := json.Unmarshal(bodyBytes, &woocommerceError)
		if err != nil {
			return ResponseDecodingError{
				Body:    bodyBytes,
				Message: err.Error(),
				Status:  r.StatusCode,
			}
		}
	}

	// Create the response error from the Shopify error.
	responseError := ResponseError{
		Status:  r.StatusCode,
		Message: woocommerceError.Error,
	}

	// If the errors field is not filled out, we can return here.
	if woocommerceError.Errors == nil {
		return wrapSpecificError(r, responseError)
	}

	switch reflect.TypeOf(woocommerceError.Errors).Kind() {
	case reflect.String:
		// Single string, use as message
		responseError.Message = woocommerceError.Errors.(string)
	case reflect.Slice:
		// An array, parse each entry as a string and join them on the message
		// json always serializes JSON arrays into []interface{}
		for _, elem := range woocommerceError.Errors.([]interface{}) {
			responseError.Data = append(responseError.Data, fmt.Sprint(elem))
		}
		responseError.Message = strings.Join(responseError.Data, ", ")
	case reflect.Map:
		// A map, parse each error for each key in the map.
		// json always serializes into map[string]interface{} for objects
		for k, v := range woocommerceError.Errors.(map[string]interface{}) {
			// Check to make sure the interface is a slice
			// json always serializes JSON arrays into []interface{}
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				for _, elem := range v.([]interface{}) {
					// If the primary message of the response error is not set, use
					// any message.
					if responseError.Message == "" {
						responseError.Message = fmt.Sprintf("%v: %v", k, elem)
					}
					topicAndElem := fmt.Sprintf("%v: %v", k, elem)
					responseError.Data = append(responseError.Data, topicAndElem)
				}
			}
		}
	}

	return wrapSpecificError(r, responseError)
}

func (c *Client) logRequest(req *http.Request) {
	if req == nil {
		return
	}
	if req.URL != nil {
		c.log.Debugf("%s: %s", req.Method, req.URL.String())
	}
	c.logBody(&req.Body, "SENT: %s")
}

func (c *Client) logResponse(res *http.Response) {
	if res == nil {
		return
	}
	c.log.Debugf("RECV %d: %s", res.StatusCode, res.Status)
	c.logBody(&res.Body, "RESP: %s")
}

func (c *Client) logBody(body *io.ReadCloser, format string) {
	if body == nil {
		return
	}
	b, _ := ioutil.ReadAll(*body)
	if len(b) > 0 {
		c.log.Debugf(format, string(b))
	}
	*body = ioutil.NopCloser(bytes.NewBuffer(b))
}

// ResponseError is A general response error that follows a similar layout to WooCommerce's response
// errors, i.e. either a single message or a list of messages.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#request-response-format
type ResponseError struct {
	Status  int
	Message string
	Data    []string
}

func (e ResponseError) Error() string {
	return e.Message
}

// An error specific to a rate-limiting response. Embeds the ResponseError to
// allow consumers to handle it the same was a normal ResponseError.
type RateLimitError struct {
	ResponseError
	RetryAfter int
}

func wrapSpecificError(r *http.Response, err ResponseError) error {
	if err.Status == http.StatusTooManyRequests {
		f, _ := strconv.ParseFloat(r.Header.Get("Retry-After"), 64)
		return RateLimitError{
			ResponseError: err,
			RetryAfter:    int(f),
		}
	}
	if err.Status == http.StatusNotAcceptable {
		err.Message = http.StatusText(err.Status)
	}

	return err
}

// CreateAndDo performs a web request to WooCommerce with the given method (GET,
// POST, PUT, DELETE) and relative path (e.g. "/wp-admin/v3").
func (c *Client) CreateAndDo(method, relPath string, data, options, resource interface{}) error {
	_, err := c.createAndDoGetHeaders(method, relPath, data, options, resource)
	if err != nil {
		return err
	}
	return nil
}

// createAndDoGetHeaders creates an executes a request while returning the response headers.
func (c *Client) createAndDoGetHeaders(method, relPath string, data, options, resource interface{}) (http.Header, error) {
	if strings.HasPrefix(relPath, "/") {
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join(c.pathPrefix, relPath)
	//println("relPath:", relPath)
	req, err := c.NewRequest(method, relPath, data, options)
	if err != nil {
		return nil, err
	}
	return c.doGetHeaders(req, resource)
}

// Creates an API request. A relative URL can be provided in urlStr, which will
// be resolved to the BaseURL of the Client. Relative URLS should always be
// specified without a preceding slash. If specified, the value pointed to by
// body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, relPath string, body, options interface{}) (*http.Request, error) {
	rel, err := url.Parse(relPath)
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// Add custom options
	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		for k, values := range u.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		u.RawQuery = optionsQuery.Encode()
	}

	// A bit of JSON ceremony
	var js []byte = nil

	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	req.SetBasicAuth(c.app.CustomerKey, c.app.CustomerSecret)
	return req, nil
}

// Get performs a GET request for the given path and saves the result in the
// given resource.
func (c *Client) Get(path string, resource, options interface{}) error {
	return c.CreateAndDo("GET", path, nil, options, resource)
}

// Post performs a POST request for the given path and saves the result in the
// given resource.
func (c *Client) Post(path string, data, resource interface{}) error {
	return c.CreateAndDo("POST", path, data, nil, resource)
}

// Put performs a PUT request for the given path and saves the result in the
// given resource.
func (c *Client) Put(path string, data, resource interface{}) error {
	return c.CreateAndDo("PUT", path, data, nil, resource)
}

// Delete performs a DELETE request for the given path
func (c *Client) Delete(path string) error {
	return c.CreateAndDo("DELETE", path, nil, nil, nil)
}

//  ListOptions represent ist options that can be used for most collections of entities.
type ListOptions struct {
	Context string  `url:"context,omitemty"`
	Page    int     `url:"page,omitemty"`
	PerPage int     `url:"per_page,omitemty"`
	Search  string  `url:"search,omitemty"`
	After   string  `url:"after,omitemty"`
	Before  string  `url:"before,omitemty"`
	Exclude []int64 `url:"exclude,omitemty"`
	Include []int64 `url:"include,omitemty"`
	Offset  int     `url:"offset,omitemty"`
	Order   string  `url:"order,omitemty"`
	Orderby string  `url:"orderby,omitemty"`
}

// OrderResource  represents the result from the /wp-json/wc/v3/orders/:id endpoint
type OrderResource struct {
	Order *Order `json:"order"`
}

// OrderResource  represents the result from the /wp-json/wc/v3/orders endpoint
type OrdersResource struct {
	Orders []Order `json:"orders"`
}
