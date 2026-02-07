package woocommerce

import (
	"testing"
	"time"
)

func TestOrderRefundServiceOp_List(t *testing.T) {
	orderID := int64(1)
	refunds, err := client.OrderRefund.List(orderID, nil)
	if err != nil {
		t.Logf("error listing refunds: %v", err)
	}
	for _, refund := range refunds {
		t.Logf("refund: id=%d, amount=%s, reason=%s", refund.ID, refund.Amount, refund.Reason)
	}
}

func TestOrderRefundServiceOp_Create(t *testing.T) {
	orderID := int64(1)
	refund := OrderRefund{
		Amount: "10.00",
		Reason: "Test refund " + time.Now().Format("20060102150405"),
	}
	res, err := client.OrderRefund.Create(orderID, refund)
	if err != nil {
		t.Logf("create refund error: %v", err)
	} else {
		t.Logf("created refund: id=%d, amount=%s", res.ID, res.Amount)
	}
}

func TestOrderRefundServiceOp_Get(t *testing.T) {
	orderID := int64(1)
	refundID := int64(1)
	refund, err := client.OrderRefund.Get(orderID, refundID, nil)
	if err != nil {
		t.Logf("get refund error: %v", err)
	} else {
		t.Logf("got refund: id=%d, amount=%s", refund.ID, refund.Amount)
	}
}

func TestOrderRefundServiceOp_Delete(t *testing.T) {
	orderID := int64(1)
	refund := OrderRefund{
		Amount: "5.00",
		Reason: "Test refund to delete " + time.Now().Format("20060102150405"),
	}
	created, err := client.OrderRefund.Create(orderID, refund)
	if err != nil {
		t.Skipf("skipping delete test: cannot create refund: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	_, err = client.OrderRefund.Delete(orderID, created.ID, optionsDel)
	if err != nil {
		t.Logf("delete refund error: %v", err)
	} else {
		t.Logf("deleted refund: id=%d", created.ID)
	}
}
