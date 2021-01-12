package woocommerce

import (
	"fmt"
	"testing"
)

const (
	customerKey    = "ck_424894d48a21577fa4b0f57394a64ae0db8e7321" // your customer_key
	customerSecret = "cs_6156c6a1e84f8337e9bad5a784d47a1a90dc7e4e" // your customer_secret
	shopUrl        = "your shop url "
)

func TestOrderServiceOp_List(t *testing.T) {
	app := App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
	}

	client := NewClient(app, shopUrl,
		WithLog(&LeveledLogger{
			Level: LevelInfo,
		}),
		WithRetry(3))
	orders, _ := client.Order.List(nil)
	for _, order := range orders {
		fmt.Println(order.ID, order.Currency)
	}
}

func TestOrderServiceOp_Get(t *testing.T) {
	app := App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
	}

	client := NewClient(app, shopUrl)
	order, err := client.Order.Get(17, nil)
	t.Logf("order : %v, err: %v", order, err)
	fmt.Println(order, err)
}
