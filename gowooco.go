package gowooco

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	defaultHttpTimeout   = 10
	defaultApiPathPrefix = "/wp-json/wc"
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

	Product   ProductService
	Order     OrderService
	OrderNote OrderNoteService
}

// NewClient returns a new WooCommerce API client with an already authenticated shopname and
// token. The shopName parameter is the shop's wooCommerce website domain,
// e.g. "shop.gitvim.com"
// a.NewClient(shopName, token, opts) is equivalent to NewClient(a, shopName, token, opts)
func (a App) NewClient(shopName string, opts ...Options) *Client {
	return NewClient(a, shopName, opts...)
}

// NewClient Returns a new WooCommerce API client with an already authenticated shopname and
// token. The shopName parameter is the shop's wooCommerce website domain,
// e.g. "shop.gitvim.com"
func NewClient(app App, shopName string, opts ...Options) *Client {
	baseURL, err := url.Parse(ShopBaseURL(shopName))
	if err != nil {
		panic(err)
	}
	c := Client{
		Client: &http.Client{
			Timeout: time.Second * defaultHttpTimeout,
		},
		log:        &LeveledLogger{},
		app:        App,
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
	return fmt.Sprintf("https://%s", name)
}
