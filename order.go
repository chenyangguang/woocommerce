package woocommerce

import (
	"fmt"
	"net/http"
)

const (
	ordersBasePath = "orders"
)

// OrderService is an interface for interfacing with the orders endpoints of woocommerce API
// https://woocommerce.github.io/woocommerce-rest-api-docs/#orders
type OrderService interface {
	Create(order Order) (*Order, error)
	Get(orderId int64, options interface{}) (*Order, error)
	List(options interface{}) ([]Order, error)
	Update(order *Order) (*Order, error)
	Delete(orderID int64, options interface{}) (*Order, error)
	Batch(option OrderBatchOption) (*OrderBatchResource, error)
}

// OrderServiceOp handles communication with the order related methods of WooCommerce'API
type OrderServiceOp struct {
	client *Client
}

// OrderListOption list all thee order list option request params
// refrence url:
// https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-orders
// parameters:
// context	string	Scope under which the request is made; determines fields present in response. Options: view and edit. Default is view.
// page	integer	Current page of the collection. Default is 1.
// per_page	integer	Maximum number of items to be returned in result set. Default is 10.
// search	string	Limit results to those matching a string.
// after	string	Limit response to resources published after a given ISO8601 compliant date.
// before	string	Limit response to resources published before a given ISO8601 compliant date.
// exclude	array	Ensure result set excludes specific IDs.
// include	array	Limit result set to specific ids.
// offset	integer	Offset the result set by a specific number of items.
// order	string	Order sort attribute ascending or descending. Options: asc and desc. Default is desc.
// orderby	string	Sort collection by object attribute. Options: date, id, include, title and slug. Default is date.
// parent	array	Limit result set to those of particular parent IDs.
// parent_exclude	array	Limit result set to all items except those of a particular parent ID.
// status	array	Limit result set to orders assigned a specific status. Options: any, pending, processing, on-hold, completed, cancelled, refunded, failed and trash. Default is any.
// customer	integer	Limit result set to orders assigned a specific customer.
// product	integer	Limit result set to orders assigned a specific product.
// dp	integer	Number of decimal points to use in each resource. Default is 2.
type OrderListOption struct {
	ListOptions
	Parent        []int64  `url:"parent,omitemty"`
	ParentExclude []int64  `url:"parent_exclude,omitemty"`
	Status        []string `url:"status,omitempty"`
	Customer      int64    `url:"customer,omitempty"`
	Product       int64    `url:"product,omitempty"`
	Dp            int      `url:"id,omitempty"`
}

// OrderDeleteOption is the only option for delete order record. dangerous
// when the force is true, it will permanently delete the order
// while the force is false, you should get the order from Get Restful API
// but the order's status became to be trash.
// it is better to setting force's column value be "false" rather then  "true"
type OrderDeleteOption struct {
	Force bool `json:"force,omitempty"`
}

// OrderBatchOption setting  operate for order in batch way
// https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-orders
type OrderBatchOption struct {
	Create []Order `json:"create,omitempty"`
	Update []Order `json:"update,omitempty"`
	Delete []int64 `json:"delete,omitempty"`
}

type OrderBatchResource struct {
	Create []*Order `json:"create,omitempty"`
	Update []*Order `json:"update,omitempty"`
	Delete []*Order `json:"delete,omitempty"`
}

// Order represents a WooCommerce Order
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-properties
type Order struct {
	ID                 int64           `json:"id,omitempty"`
	ParentId           int64           `json:"parent_id,omitempty"`
	Number             string          `json:"number,omitempty"`
	OrderKey           string          `json:"order_key,omitempty"`
	CreatedVia         string          `json:"created_via,omitempty"`
	Version            string          `json:"version,omitempty"`
	Status             string          `json:"status,omitempty"`
	Currency           string          `json:"currency,omitempty"`
	DateCreated        string          `json:"date_created,omitempty"`
	DateCreatedGmt     string          `json:"date_created_gmt,omitempty"`
	DateModified       string          `json:"date_modified,omitempty"`
	DateModifiedGmt    string          `json:"date_modified_gmt,omitempty"`
	DiscountsTotal     string          `json:"discount_total,omitempty"`
	DiscountsTax       string          `json:"discount_tax,omitempty"`
	ShippingTotal      string          `json:"shipping_total,omitempty"`
	ShippingTax        string          `json:"shipping_tax,omitempty"`
	CartTax            string          `json:"cart_tax,omitempty"`
	Total              string          `json:"total,omitempty"`
	TotalTax           string          `json:"total_tax,omitempty"`
	PricesIncludeTax   bool            `json:"prices_include_tax,omitempty"`
	CustomerId         int64           `json:"customer_id,omitempty"`
	CustomerIpAddress  string          `json:"customer_ip_address,omitempty"`
	CustomerUserAgent  string          `json:"customer_user_agent,omitempty"`
	CustomerNote       string          `json:"customer_note,omitempty"`
	Billing            *Billing        `json:"billing,omitempty"`
	Shipping           *Shipping       `json:"shipping,omitempty"`
	PaymentMethod      string          `json:"payment_method,omitempty"`
	PaymentMethodTitle string          `json:"payment_method_title,omitempty"`
	TransactionId      string          `json:"transaction_id,omitempty"`
	DatePaid           string          `json:"date_paid,omitempty"`
	DatePaidGmt        string          `json:"date_paid_gmt,omitempty"`
	DateCompleted      string          `json:"date_completed,omitempty"`
	DateCompletedGmt   string          `json:"date_completed_gmt,omitempty"`
	CartHash           string          `json:"cart_hash,omitempty"`
	MetaData           []MetaData      `json:"meta_data,omitempty"`
	LineItems          []LineItem      `json:"line_items,omitempty"`
	TaxLines           []TaxLine       `json:"tax_lines,omitempty"`
	ShippingLines      []ShippingLines `json:"shipping_lines,omitempty"`
	FeeLines           []FeeLine       `json:"fee_lines,omitempty"`
	CouponLines        []CouponLine    `json:"coupon_lines,omitempty"`
	Refunds            []Refund        `json:"refunds,omitempty"`
	CurrencySymbol     string          `json:"currency_symbol,omitempty"`
	Links              Links           `json:"_links"`
	SetPaid            bool            `json:"set_paid,omitempty"`
}

type Links struct {
	Self []struct {
		Href string `json:"href"`
	} `json:"self"`
	Collection []struct {
		Href string `json:"href"`
	} `json:"collection"`
}

type Billing struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"province,omitempty"`
	PostCode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type Shipping struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"province,omitempty"`
	PostCode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
}

type LineItem struct {
	ID          int64      `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	ProductID   int64      `json:"product_id,omitempty"`
	VariantID   int64      `json:"variant_id,omitempty"`
	Quantity    int        `json:"quantity,omitempty"`
	TaxClass    string     `json:"tax_class,omitempty"`
	SubTotal    string     `json:"subtotal,omitempty"`
	SubtotalTax string     `json:"subtotal_tax,omitempty"`
	Total       string     `json:"total,omitempty"`
	TotalTax    string     `json:"total_tax,omitempty"`
	Taxes       []TaxLine  `json:"taxes,omitempty"`
	MetaData    []MetaData `json:"meta_data,omitempty"`
	SKU         string     `json:"sku,omitempty"`
	Price       int64      `json:"price,omitempty"`
}

type TaxLine struct {
	ID               int64      `json:"id,omitempty"`
	RateCode         string     `json:"rate_code,omitempty"`
	RateId           string     `json:"rate_id,omitempty"`
	Label            string     `json:"label,omitempty"`
	Compound         bool       `json:"compound,omitempty"`
	TaxTotal         string     `json:"tax_total"`
	ShippingTaxTotal string     `json:"shipping_tax_total,omitempty"`
	MetaData         []MetaData `json:"meta_data,omitempty"`
}

type MetaData struct {
	ID    int64  `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type FeeLine struct {
	ID        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	TaxClass  string     `json:"tax_class,omitempty"`
	TaxStatus string     `json:"tax_status,omitempty"`
	Total     string     `json:"total,omitempty"`
	TotalTax  string     `json:"total_tax,omitempty"`
	Taxes     []TaxLine  `json:"taxes,omitempty"`
	MetaData  []MetaData `json:"meta_data,omitempty"`
}

type Refund struct {
	ID     int64  `json:"id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Total  string `json:"total,omitempty"`
}

type ShippingLines struct {
	ID          int64      `json:"id,omitempty"`
	MethodTitle string     `json:"method_title,omitempty"`
	MethodID    string     `json:"method_id,omitempty"`
	Total       string     `json:"total,omitempty"`
	TotalTax    string     `json:"total_tax,omitempty"`
	Taxes       []TaxLine  `json:"tax_lines,omitempty"`
	MetaData    []MetaData `json:"meta_data,omitempty"`
}

type CouponLine struct {
	ID          int64      `json:"id,omitempty"`
	Code        string     `json:"code,omitempty"`
	Discount    string     `json:"discount,omitempty"`
	DiscountTax string     `json:"discount_tax,omitempty"`
	MetaData    []MetaData `json:"meta_data,omitempty"`
}

func (o *OrderServiceOp) List(options interface{}) ([]Order, error) {
	orders, _, err := o.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (o *OrderServiceOp) ListWithPagination(options interface{}) ([]Order, *Pagination, error) {
	path := fmt.Sprintf("%s", ordersBasePath)
	resource := make([]Order, 0)
	headers := http.Header{}
	headers, err := o.client.createAndDoGetHeaders("GET", path, nil, options, &resource)
	if err != nil {
		return nil, nil, err
	}
	// Extract pagination info from header
	linkHeader := headers.Get("Link")
	println(linkHeader)
	//pagination, err := extractPagination(linkHeader)
	//if err != nil {
	//	return nil, nil, err
	//}

	return resource, nil, nil
}

func (o *OrderServiceOp) Create(order Order) (*Order, error) {
	path := fmt.Sprintf("%s", ordersBasePath)
	resource := new(Order)

	err := o.client.Post(path, order, &resource)
	return resource, err
}

// Get individual order
func (o *OrderServiceOp) Get(orderID int64, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d", ordersBasePath, orderID)
	resource := new(Order)
	err := o.client.Get(path, resource, options)
	return resource, err
}

func (o *OrderServiceOp) Update(order *Order) (*Order, error) {
	path := fmt.Sprintf("%s/%d", ordersBasePath, order.ID)
	resource := new(Order)
	err := o.client.Put(path, order, &resource)
	return resource, err
}

func (o *OrderServiceOp) Delete(orderID int64, options interface{}) (*Order, error) {
	path := fmt.Sprintf("%s/%d", ordersBasePath, orderID)
	resource := new(Order)
	err := o.client.Delete(path, options, &resource)
	return resource, err
}

func (o *OrderServiceOp) Batch(data OrderBatchOption) (*OrderBatchResource, error) {
	path := fmt.Sprintf("%s/batch", ordersBasePath)
	resource := new(OrderBatchResource)
	err := o.client.Post(path, data, &resource)
	return resource, err
}
