package woocommerce

import "testing"

func TestPaymentGatewayServiceOp_List(t *testing.T) {
	payments, err := client.PaymentGateway.List(nil)
	if err != nil {
		t.Errorf("get payment list fail: %v", err)
		t.FailNow()
	}
	for _, payment := range payments {
		t.Log(payment.ID, payment.Title, payment)
	}
}

func TestPaymentGatewayServiceOp_Get(t *testing.T) {
	payment, err := client.PaymentGateway.Get("paypal")
	if err != nil {
		t.Errorf("get payment fail: %v", err)
		t.FailNow()
	}
	t.Logf("payment : %v, err: %v", payment.Settings.Email.Value, err)
}
