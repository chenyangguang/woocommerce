package woocommerce

import (
	"fmt"
)

const (
	orderRefundBasePath = "orders"
)

// OrderRefundService allows you to create, view, and delete individual WooCommerce Order refunds.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-refunds
type OrderRefundService interface {
	Create(orderID int64, refund OrderRefund) (*OrderRefund, error)
	Get(orderID int64, refundID int64, options interface{}) (*OrderRefund, error)
	List(orderID int64, options interface{}) ([]OrderRefund, error)
	Delete(orderID int64, refundID int64, options interface{}) (*OrderRefund, error)
}

// OrderRefund represent a WooCommerce Order Refund
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-refund-properties
type OrderRefund struct {
	ID             int64  `json:"id,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateCreatedGmt string `json:"date_created_gmt,omitempty"`
	Amount         string `json:"amount,omitempty"`
	Reason         string `json:"reason,omitempty"`
	RefundedBy     int64  `json:"refunded_by,omitempty"`
}

type OrderRefundServiceOp struct {
	client *Client
}

func (o *OrderRefundServiceOp) Create(orderID int64, refund OrderRefund) (*OrderRefund, error) {
	path := fmt.Sprintf("%s/%d/refunds", orderRefundBasePath, orderID)
	resource := new(OrderRefund)
	err := o.client.Post(path, refund, &resource)
	return resource, err
}

func (o *OrderRefundServiceOp) Get(orderID int64, refundID int64, options interface{}) (*OrderRefund, error) {
	path := fmt.Sprintf("%s/%d/refunds/%d", orderRefundBasePath, orderID, refundID)
	resource := new(OrderRefund)
	err := o.client.Get(path, resource, options)
	return resource, err
}

func (o *OrderRefundServiceOp) List(orderID int64, options interface{}) ([]OrderRefund, error) {
	path := fmt.Sprintf("%s/%d/refunds", orderRefundBasePath, orderID)
	resource := make([]OrderRefund, 0)
	err := o.client.Get(path, &resource, options)
	return resource, err
}

func (o *OrderRefundServiceOp) Delete(orderID int64, refundID int64, options interface{}) (*OrderRefund, error) {
	path := fmt.Sprintf("%s/%d/refunds/%d", orderRefundBasePath, orderID, refundID)
	resource := new(OrderRefund)
	err := o.client.Delete(path, options, &resource)
	return resource, err
}
