package woocommerce

import (
	"fmt"
	"testing"
	"time"
)

const (
	customerKey    = "customer_key"    // your customer_key
	customerSecret = "customer_secret" // your customer_secret
	shopUrl        = "shop.gitvim.com" // your shop website domain
)

var client *Client

func init() {
	app := App{
		CustomerKey:    customerKey,
		CustomerSecret: customerSecret,
	}

	client = NewClient(app, shopUrl,
		WithLog(&LeveledLogger{
			Level: LevelDebug, // you should open this for debug in dev environment,  usefully.
		}),
		WithRetry(3))
}
func TestOrderServiceOp_List(t *testing.T) {
	options := OrderListOption{
		ListOptions: ListOptions{
			Context: "view",
			After:   "2021-01-01T06:16:17",
			Before:  "2022-01-12T06:16:17",
			Order:   "desc",
			Orderby: "date",
			Page:    2,
			PerPage: 2,
		},
		Status:  []string{"processing"},
		Product: 10,
	}
	orders, err := client.Order.List(options)
	if err != nil {
		fmt.Println("err result: ", err)
	}
	for _, order := range orders {
		t.Log(order.ID, order.Currency)
	}
}

func TestOrderServiceOp_Get(t *testing.T) {
	order, err := client.Order.Get(17, nil)
	t.Logf("order : %v, err: %v", order, err)
}

func initOrder() Order {
	timeNow := time.Now().Unix()
	timeNowStr := fmt.Sprintf("%d", timeNow)
	order := Order{
		PaymentMethod: "paypal",
		Billing: &Billing{
			FirstName: "git" + timeNowStr,
			LastName:  "vim" + timeNowStr,
		},
		LineItems: []LineItem{
			{
				Name:      "北京烤鸭" + timeNowStr,
				ProductID: 10,
				SubTotal:  "56.00",
				Total:     "56.00",
				Quantity:  2,
				MetaData: []MetaData{
					{
						Key:   "_reduced_stock",
						Value: "2",
					},
				},
				SKU:   "wutongshan_001" + timeNowStr,
				Price: 56.00,
			},
		},
	}
	return order
}

func TestOrderServiceOp_Create(t *testing.T) {
	order := initOrder()
	res, err := client.Order.Create(order)
	if err != nil {
		t.Errorf("res : %v, err: %v", res, err)
	} else {
		t.Logf("res:%v", res)
	}
}

func TestOrderServiceOp_Update(t *testing.T) {
	order, err := client.Order.Get(17, nil)
	if order == nil || err != nil {
		t.Errorf("get order fail : %v", err)
	}
	order.Currency = "CNY"
	res, err := client.Order.Update(order)
	if err != nil {
		t.Errorf("update order fail: %v", err)
	}
	t.Logf("update success result is : %v", res)
}

func TestOrderServiceOp_Delete(t *testing.T) {
	optionsDel := DeleteOption{
		Force: false,
	}
	res, err := client.Order.Delete(29, optionsDel)
	if err != nil {
		t.Errorf("delete order fail: %v", err)
		t.FailNow()
	}
	t.Logf("delete order result : %v", res)
	// go test -v -run=TestOrderServiceOp_Delete
}

func TestOrderServiceOp_Batch(t *testing.T) {
	order := initOrder()
	data := OrderBatchOption{
		Create: []Order{
			order,
		},
		Update: []Order{
			{
				ID:       17,
				TotalTax: "20.00",
				Total:    "120",
			},
		},
		Delete: []int64{
			18,
		},
	}
	res, err := client.Order.Batch(data)
	if err != nil {
		t.Errorf("delete order fail: %v", err)
		t.FailNow()
	}
	t.Logf(" create : %v, update: %v, delete : %v", res.Create, res.Update, res.Delete)
	for _, order := range res.Create {
		t.Logf(" order id: %v, order total : %v", order.ID, order.Total)
	}

	for _, order := range res.Update {
		t.Logf(" order id: %v, order total : %v", order.ID, order.Total)
	}

	for _, order := range res.Delete {
		t.Logf(" order id: %v, order status : %v", order.ID, order.Status)
	}
}
