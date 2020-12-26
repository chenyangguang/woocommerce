package gowooco

import (
	"github.com/shopspring/decimal"
	"time"
)

// OrderRefundService allows you to create, view, and delete individual WooCommerce Order refunds.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#refunds
type OrderRefundService interface {
	Create()
	Get()
	Delete()
	List()
}

// OrderRefund represent a WooCommerce Order Refund
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-refund-properties
type OrderRefund struct {
	ID int64 `json:"id,omitempty"`

	DateCreated    *time.Time `json:"date_created,omitempty"`
	DateCreatedGmt *time.Time `json:"date_created_gmt,omitempty"`

	Amount *decimal.Decimal `json:"amount,omitempty"`
}
