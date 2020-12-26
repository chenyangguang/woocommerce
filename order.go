package gowc

import (
	"github.com/shopspring/decimal"
	"time"
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
