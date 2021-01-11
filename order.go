package gowooco

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const (
	baseOrderPath      = "order"
	baseOrdersResource = "orders"
)

// OrderService is an interface for interfacing with the orders endpoints of woocommerce API
// https://woocommerce.github.io/woocommerce-rest-api-docs/#orders
type OrderService interface {
	Create()
	Get()
	List()
	Update()
	Delete()
	BatchUpdate()
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
	ListOption
	Parent        []int64  `url:"parent,omitemty"`
	ParentExclude []int64  `url:"parent_exclude,omitemty"`
	Status        []string `url:"status,omitempty"`
	Customer      int64    `url:"customer,omitempty"`
	Product       int64    `url:"product,omitempty"`
	Dp            int      `url:"id,omitempty"`
}

// Order represents a WooCommerce Order
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-properties
type Order struct {
	ID                 int64            `json:"id,omitempty"`
	ParentId           int64            `json:"parent_id,omitempty"`
	Number             string           `json:"number,omitempty"`
	OrderKey           string           `json:"order_key,omitempty"`
	CreatedVia         string           `json:"created_via,omitempty"`
	Version            string           `json:"version,omitempty"`
	Status             string           `json:"status,omitempty"`
	Currency           string           `json:"currency,omitempty"`
	DateCreated        *time.Time       `json:"date_created,omitempty"`
	DateCreatedGmt     *time.Time       `json:"date_created_gmt,omitempty"`
	DateModified       *time.Time       `json:"date_modified,omitempty"`
	DateModifiedGmt    *time.Time       `json:"date_modified_gmt,omitempty"`
	DiscountsTotal     *decimal.Decimal `json:"discount_total,omitempty"`
	DiscountsTax       *decimal.Decimal `json:"discount_tax,omitempty"`
	ShippingTotal      *decimal.Decimal `json:"shipping_total,omitempty"`
	ShippingTax        *decimal.Decimal `json:"shipping_tax,omitempty"`
	CartTax            *decimal.Decimal `json:"cart_tax,omitempty"`
	Total              *decimal.Decimal `json:"total,omitempty"`
	TotalTax           *decimal.Decimal `json:"total_tax,omitempty"`
	PricesIncludeTax   bool             `json:"prices_include_tax,omitempty"`
	CustomerId         int64            `json:"customer_id,omitempty"`
	CustomerIpAddress  string           `json:"customer_ip_address,omitempty"`
	CustomerUserAgent  string           `json:"customer_user_agent,omitempty"`
	CustomerNote       string           `json:"customer_note,omitempty"`
	Billing            *Address         `json:"billing,omitempty"`
	Shipping           *Address         `json:"shipping,omitempty"`
	PaymentMethod      string           `json:"payment_method,omitempty"`
	PaymentMethodTitle string           `json:"payment_method_title,omitempty"`
	TransactionId      string           `json:"transaction_id,omitempty"`
	DatePaid           time.Time        `json:"date_paid,omitempty"`
	DatePaidGmt        time.Time        `json:"date_paid_gmt,omitempty"`
	DateCompleted      time.Time        `json:"date_completed,omitempty"`
	DateCompletedGmt   time.Time        `json:"date_completed_gmt,omitempty"`
	CartHash           string           `json:"cart_hash,omitempty"`
	MetaData           []MetaData       `json:"meta_data,omitempty"`
	LineItems          []LineItem       `json:"line_items,omitempty"`
	TaxLines           []TaxLine        `json:"tax_lines,omitempty"`
	ShippingLines      []ShippingLines  `json:"shipping_lines,omitempty"`
	FeeLines           []FeeLine        `json:"fee_lines,omitempty"`
	CouponLines        []CouponLine     `json:"coupon_lines,omitempty"`
	Refunds            []Refund         `json:"refunds,omitempty"`
	SetPaid            bool             `json:"set_paid,omitempty"`
}

type Address struct {
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

type LineItem struct {
	ID          int64            `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	ProductID   int64            `json:"product_id,omitempty"`
	VariantID   int64            `json:"variant_id,omitempty"`
	Quantity    int              `json:"quantity,omitempty"`
	TaxClass    string           `json:"tax_class,omitempty"`
	SubTotal    *decimal.Decimal `json:"subtotal,omitempty"`
	SubtotalTax *decimal.Decimal `json:"subtotal_tax,omitempty"`
	Total       *decimal.Decimal `json:"total,omitempty"`
	TotalTax    *decimal.Decimal `json:"total_tax,omitempty"`
	Taxes       []TaxLine        `json:"taxes,omitempty"`
	MetaData    []MetaData       `json:"meta_data,omitempty"`
	SKU         string           `json:"sku,omitempty"`
	Price       *decimal.Decimal `json:"price,omitempty"`
}

type TaxLine struct {
	ID               int64            `json:"id,omitempty"`
	RateCode         string           `json:"rate_code,omitempty"`
	RateId           string           `json:"rate_id,omitempty"`
	Label            string           `json:"label,omitempty"`
	Compound         bool             `json:"compound,omitempty"`
	TaxTotal         *decimal.Decimal `json:"tax_total"`
	ShippingTaxTotal *decimal.Decimal `json:"shipping_tax_total,omitempty"`
	MetaData         []MetaData       `json:"meta_data,omitempty"`
}

type MetaData struct {
	ID    int64  `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type FeeLine struct {
	ID        int64            `json:"id,omitempty"`
	Name      string           `json:"name,omitempty"`
	TaxClass  string           `json:"tax_class,omitempty"`
	TaxStatus string           `json:"tax_status,omitempty"`
	Total     *decimal.Decimal `json:"total,omitempty"`
	TotalTax  *decimal.Decimal `json:"total_tax,omitempty"`
	Taxes     []TaxLine        `json:"taxes,omitempty"`
	MetaData  []MetaData       `json:"meta_data,omitempty"`
}

type Refund struct {
	ID     int64            `json:"id,omitempty"`
	Reason string           `json:"reason,omitempty"`
	Total  *decimal.Decimal `json:"total,omitempty"`
}

type ShippingLines struct {
	ID          int64            `json:"id,omitempty"`
	MethodTitle string           `json:"method_title,omitempty"`
	MethodID    string           `json:"method_id,omitempty"`
	Total       *decimal.Decimal `json:"total,omitempty"`
	TotalTax    *decimal.Decimal `json:"total_tax,omitempty"`
	Taxes       []TaxLine        `json:"tax_lines,omitempty"`
	MetaData    []MetaData       `json:"meta_data,omitempty"`
}

type CouponLine struct {
	ID          int64            `json:"id,omitempty"`
	Code        string           `json:"code,omitempty"`
	Discount    *decimal.Decimal `json:"discount,omitempty"`
	DiscountTax *decimal.Decimal `json:"discount_tax,omitempty"`
	MetaData    []MetaData       `json:"meta_data,omitempty"`
}

func (o *OrderServiceOp) List(option interface{}) ([]order, error) {
	orders, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrderServiceOp) ListWithPagination(options interface{}) ([]order, error) {
	path := fmt.Sprintf("%s", orderBasePath)
	resource := new(OrderResource)
	headers := http.Header{}
	headers, err := o.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, err
	}
	return resource.Orders, nil
}
