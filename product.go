package woocommerce

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	productsBasePath = "products"
)

var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(prev|next|first|last)" *$`)

// ProductService allows you to create, view, update, and delete individual, or a batch, of products
// https://woocommerce.github.io/woocommerce-rest-api-docs/#products
type ProductService interface {
	Create(product Product) (*Product, error)
	Get(productID int64, options interface{}) (*Product, error)
	List(options interface{}) ([]Product, error)
	Update(product *Product) (*Product, error)
	Delete(productID int64, options interface{}) (*Product, error)
	Batch(data ProductBatchOption) (*ProductBatchResource, error)
}

// Product represent WooCommerce Product
// https://woocommerce.github.io/woocommerce-rest-api-docs/#product-properties
type Product struct {
	ID                int64                  `json:"id,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Slug              string                 `json:"slug,omitempty"`
	Permalink         string                 `json:"permalink,omitempty"`
	DateCreated       string                 `json:"date_created,omitempty"`
	DateCreatedGmt    string                 `json:"date_created_gmt,omitempty"`
	DateModified      string                 `json:"date_modified,omitempty"`
	DateModifiedGmt   string                 `json:"date_modified_gmt,omitempty"`
	Type              string                 `json:"type,omitempty"`
	Status            string                 `json:"status,omitempty"`
	Featured          bool                   `json:"featured,omitempty"`
	CatalogVisibility string                 `json:"catalog_visibility,omitempty"`
	Description       string                 `json:"description,omitempty"`
	ShortDescription  string                 `json:"short_description,omitempty"`
	SKU               string                 `json:"sku,omitempty"`
	Price             string                 `json:"price,omitempty"`
	RegularPrice      string                 `json:"regular_price,omitempty"`
	SalePrice         string                 `json:"sale_price,omitempty"`
	DateOnSaleFrom    string                 `json:"date_on_sale_from,omitempty"`
	DateOnSaleTo      string                 `json:"date_on_sale_to,omitempty"`
	TotalSales        int64                  `json:"total_sales,omitempty"`
	TaxStatus         string                 `json:"tax_status,omitempty"`
	TaxClass          string                 `json:"tax_class,omitempty"`
	ManageStock       bool                   `json:"manage_stock,omitempty"`
	StockQuantity     string                 `json:"stock_quantity,omitempty"`
	StockStatus       string                 `json:"stock_status,omitempty"`
	Backorders        string                 `json:"backorders,omitempty"`
	LowStockAmount    string                 `json:"low_stock_amount,omitempty"`
	SoldIndividually  bool                   `json:"sold_individually,omitempty"`
	Weight            string                 `json:"weight,omitempty"`
	Length            string                 `json:"length,omitempty"`
	Width             string                 `json:"width,omitempty"`
	Height            string                 `json:"height,omitempty"`
	Dimensions        map[string]string      `json:"dimensions,omitempty"`
	ShippingClass     string                 `json:"shipping_class,omitempty"`
	ShippingRequired  bool                   `json:"shipping_required,omitempty"`
	ShippingTaxable   bool                   `json:"shipping_taxable,omitempty"`
	ShippingClassID   int64                  `json:"shipping_class_id,omitempty"`
	ReviewsAllowed    string                 `json:"reviews_allowed,omitempty"`
	AverageRating     string                 `json:"average_rating,omitempty"`
	RatingCount       int64                  `json:"rating_count,omitempty"`
	RelatedIDs        []int64                `json:"related_ids,omitempty"`
	UpsellIDs         []int64                `json:"upsell_ids,omitempty"`
	CrossSellIDs      []int64                `json:"cross_sell_ids,omitempty"`
	ParentID          int64                  `json:"parent_id,omitempty"`
	PurchaseNote      string                 `json:"purchase_note,omitempty"`
	MenuOrder         int                    `json:"menu_order,omitempty"`
	Attributes        []ProductAttribute     `json:"attributes,omitempty"`
	DefaultAttributes map[string]interface{} `json:"default_attributes,omitempty"`
	Categories        []int64                `json:"categories,omitempty"`
	Tags              []int64                `json:"tags,omitempty"`
	Images            []ProductImage         `json:"images,omitempty"`
	Downloads         []ProductDownload      `json:"downloads,omitempty"`
	MetaData          []MetaData             `json:"meta_data,omitempty"`
	Links             Links                  `json:"_links,omitempty"`
}

// ProductListOption list all the product list option request params
type ProductListOption struct {
	ListOptions
	Search        string  `url:"search,omitempty"`
	After         string  `url:"after,omitempty"`
	Before        string  `url:"before,omitempty"`
	Exclude       []int64 `url:"exclude,omitempty"`
	Include       []int64 `url:"include,omitempty"`
	Offset        int     `url:"offset,omitempty"`
	Order         string  `url:"order,omitempty"`
	Orderby       string  `url:"orderby,omitempty"`
	Type          string  `url:"type,omitempty"`
	SKU           string  `url:"sku,omitempty"`
	Featured      bool    `url:"featured,omitempty"`
	Category      []int64 `url:"category,omitempty"`
	Tag           []int64 `url:"tag,omitempty"`
	ShippingClass string  `url:"shipping_class,omitempty"`
	Attribute     string  `url:"attribute,omitempty"`
	AttributeTerm string  `url:"attribute_term,omitempty"`
	StockStatus   string  `url:"stock_status,omitempty"`
}

// ProductBatchOption setting  operate for product in batch way
type ProductBatchOption struct {
	Create []Product `json:"create,omitempty"`
	Update []Product `json:"update,omitempty"`
	Delete []int64   `json:"delete,omitempty"`
}

// ProductBatchResource conservation the response struct for ProductBatchOption request
type ProductBatchResource struct {
	Create []*Product `json:"create,omitempty"`
	Update []*Product `json:"update,omitempty"`
	Delete []*Product `json:"delete,omitempty"`
}

type ProductAttribute struct {
	ID        int64                    `json:"id,omitempty"`
	Name      string                   `json:"name,omitempty"`
	Slug      string                   `json:"slug,omitempty"`
	Position  int                      `json:"position,omitempty"`
	Visible   bool                     `json:"visible,omitempty"`
	Variation bool                     `json:"variation,omitempty"`
	Options   []ProductAttributeOption `json:"options,omitempty"`
}

type ProductReview struct {
	ID            int64  `json:"id,omitempty"`
	ProductID     int64  `json:"product_id,omitempty"`
	Rating        int64  `json:"rating,omitempty"`
	Review        string `json:"review,omitempty"`
	Reviewer      string `json:"reviewer,omitempty"`
	ReviewerEmail string `json:"reviewer_email,omitempty"`
	DateCreated   string `json:"date_created,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Hold          bool   `json:"hold,omitempty"`
}

type ProductAttributeOption struct {
	Name   string `json:"name,omitempty"`
	Slug   string `json:"slug,omitempty"`
	Option string `json:"option,omitempty"`
}

type ProductImage struct {
	ID              int64  `json:"id,omitempty"`
	DateCreated     string `json:"date_created,omitempty"`
	DateCreatedGmt  string `json:"date_created_gmt,omitempty"`
	DateModified    string `json:"date_modified,omitempty"`
	DateModifiedGmt string `json:"date_modified_gmt,omitempty"`
	Src             string `json:"src,omitempty"`
	Name            string `json:"name,omitempty"`
	Alt             string `json:"alt,omitempty"`
	Position        int64  `json:"position,omitempty"`
}

type ProductDownload struct {
	ID             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Slug           string `json:"slug,omitempty"`
	File           string `json:"file,omitempty"`
	FilePath       string `json:"file_path,omitempty"`
	DownloadLimit  int64  `json:"download_limit,omitempty"`
	DownloadExpiry int64  `json:"download_expiry,omitempty"`
	DownloadType   string `json:"download_type,omitempty"`
	DownloadCount  int64  `json:"download_count,omitempty"`
}

type ProductServiceOp struct {
	client *Client
}

// Pagination of results
type Pagination struct {
	NextPageOptions     *ListOptions
	PreviousPageOptions *ListOptions
	FirstPageOptions    *ListOptions
	LastPageOptions     *ListOptions
}

func (p *ProductServiceOp) List(options interface{}) ([]Product, error) {
	path := fmt.Sprintf("%s", productsBasePath)
	resource := make([]Product, 0)
	err := p.client.Get(path, &resource, options)
	return resource, err
}

func (p *ProductServiceOp) Create(product Product) (*Product, error) {
	path := fmt.Sprintf("%s", productsBasePath)
	resource := new(Product)
	err := p.client.Post(path, product, &resource)
	return resource, err
}

// Get individual product
func (p *ProductServiceOp) Get(productID int64, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%d", productsBasePath, productID)
	resource := new(Product)
	err := p.client.Get(path, resource, options)
	return resource, err
}

func (p *ProductServiceOp) Update(product *Product) (*Product, error) {
	path := fmt.Sprintf("%s/%d", productsBasePath, product.ID)
	resource := new(Product)
	err := p.client.Put(path, product, &resource)
	return resource, err
}

func (p *ProductServiceOp) Delete(productID int64, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%d", productsBasePath, productID)
	resource := new(Product)
	err := p.client.Delete(path, options, &resource)
	return resource, err
}

func (p *ProductServiceOp) Batch(data ProductBatchOption) (*ProductBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productsBasePath)
	resource := new(ProductBatchResource)
	err := p.client.Post(path, data, &resource)
	return resource, err
}

// extractPagination extracts pagination info from linkHeader.
// Details on the format are here:
// https://woocommerce.github.io/woocommerce-rest-api-docs/#pagination
// Link: <https://www.example.com/wp-json/wc/v3/products?page=2>; rel="next",
// <https://www.example.com/wp-json/wc/v3/products?page=3>; rel="last"`
func extractPagination(linkHeader string) (*Pagination, error) {
	pagination := new(Pagination)

	if linkHeader == "" {
		return pagination, nil
	}

	for _, link := range strings.Split(linkHeader, ",") {
		match := linkRegex.FindStringSubmatch(link)
		// Make sure the link is not empty or invalid
		if len(match) != 4 {
			// We expect 3 values:
			// match[0] = full match
			// match[1] is the URL and match[2] is either 'previous' or 'next', 'first', 'last'
			err := ResponseDecodingError{
				Message: "could not extract pagination link header",
			}
			return nil, err
		}

		rel, err := url.Parse(match[1])
		if err != nil {
			err = ResponseDecodingError{
				Message: "pagination does not contain a valid URL",
			}
			return nil, err
		}

		params, err := url.ParseQuery(rel.RawQuery)
		if err != nil {
			return nil, err
		}

		paginationListOptions := ListOptions{}

		page := params.Get("page")
		if page != "" {
			paginationListOptions.Page, err = strconv.Atoi(params.Get("page"))
			if err != nil {
				return nil, err
			}
		}

		switch match[2] {
		case "next":
			pagination.NextPageOptions = &paginationListOptions
		case "prev":
			pagination.PreviousPageOptions = &paginationListOptions
		case "first":
			pagination.FirstPageOptions = &paginationListOptions
		case "last":
			pagination.LastPageOptions = &paginationListOptions
		}

	}

	return pagination, nil
}
