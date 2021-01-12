package woocommerce

import (
	"fmt"
	"testing"
)

const (
	customerKey    = "ck_424894d48a21577fa4b0f57394a64ae0db8e7321" // your customer_key
	customerSecret = "cs_6156c6a1e84f8337e9bad5a784d47a1a90dc7e4e" // your customer_secret
	shopUrl        = "testwcorder.salinshop.com"
)
var client  *Client
func init () {
	app := App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
	}

	client = NewClient(app, shopUrl,
		WithLog(&LeveledLogger{
			Level: LevelDebug,
		}),
		WithRetry(3))
}
func TestOrderServiceOp_List(t *testing.T) {

	orders, _ := client.Order.List(nil)
	for _, order := range orders {
		println(order.ID)
		fmt.Println(order.ID, order.Currency)
		t.Log(order.ID,order)
	}
}

func TestOrderServiceOp_Get(t *testing.T) {
	order, err := client.Order.Get(17, nil)
	t.Logf("order : %v, err: %v", order, err)
	fmt.Println(order, err)
}

func TestOrderServiceOp_Create(t *testing.T) {
	order := Order{
		PaymentMethod:  "paypal",
		Billing: &Billing{
			FirstName: "git",
			LastName: "vim",
		},
		LineItems: []LineItem{{
			ProductID: 93,
			Quantity: 2,
		},
		},
	}
	res, err := client.Order.Create(order)
	fmt.Println(res, err)
	if err != nil {
		t.Errorf("res : %v, err: %v", res, err)
	} else {
		t.Logf("res:%v", res)
	}
}
