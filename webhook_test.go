package woocommerce

import (
	"fmt"
	"testing"
	"time"
)

func initWebhook() Webhook {
	timeNowStr := fmt.Sprintf("%d", time.Now().Unix())
	webhook := Webhook{
		Name:        "order create" + timeNowStr,
		Topic:       "order.created",
		DeliveryUrl: "https://shop.gitvim.com", // your callback url for wooCommerce event cron job to notify
	}
	return webhook
}
func TestWebhookServiceOp_List(t *testing.T) {
	webhooks, err := client.Webhook.List(nil)
	if err != nil {
		t.Errorf("get webhook fail: %v", err)
		t.FailNow()
	}
	for _, webhook := range webhooks {
		t.Log(webhook.ID, webhook.Topic, webhook.Event, webhook.Hooks)
		// TestWebhookServiceOp_List: webhook_test.go:12: 5 order.updated updated [woocommerce_update_order woocommerce_order_refunded]
		//	TestWebhookServiceOp_List: webhook_test.go:12: 4 order.updated updated [woocommerce_update_order woocommerce_order_refunded]
		//	TestWebhookServiceOp_List: webhook_test.go:12: 3 order.updated updated [woocommerce_update_order woocommerce_order_refunded]
		// ...
	}
}

func TestWebhookServiceOp_Create(t *testing.T) {
	webhook := initWebhook()
	res, err := client.Webhook.Create(webhook)
	if err != nil {
		t.Errorf("res : %v, err: %v", res, err)
		t.FailNow()
	} else {
		t.Logf("res:%v", res)
	}
}

func TestWebhookServiceOp_Get(t *testing.T) {
	webhook, err := client.Webhook.Get(2, nil)
	if err != nil {
		t.Errorf("get webhook fail: %v", err)
		t.FailNow()
	}
	t.Logf("order : %v, err: %v", webhook, err)
}

func TestWebhookServiceOp_Update(t *testing.T) {
	webhook, err := client.Webhook.Get(2, nil)
	if err != nil {
		return
	}
	webhook.Name = webhook.Name + " after updated"
	res, err := client.Webhook.Update(webhook)
	if err != nil {
		t.Errorf("update webhook fail: %v", err)
		t.FailNow()
	}
	t.Logf("webhook result: %v", res)
}

func TestWebhookServiceOp_Delete(t *testing.T) {
	options := DeleteOption{
		Force: true,
	}
	webhook, err := client.Webhook.Delete(9, options)
	if err != nil {
		t.Errorf("delete webhook fail: %v", err)
		// if you don't set the options for force = true, you will get the result like below:
		// {"code":"woocommerce_rest_trash_not_supported","message":"Webhook\u4e0d\u652f\u6301\u56de\u6536\u7ad9\u3002","data":{"status":501}}
		// But delete a webhook is very dangerous operation
		// if you absolute need to delete a webhook , just setting force's value as true
	}
	t.Logf("The webhook you delete is : %v", webhook)
}

func TestWebhookServiceOp_Batch(t *testing.T) {
	webhook := initWebhook()
	data := WebhookBatchOption{
		Create: []Webhook{
			webhook,
		},
		Update: []Webhook{
			{
				ID:   8,
				Name: "batch update operate test",
			},
		},
		Delete: []int64{
			10,
		},
	}
	res, err := client.Webhook.Batch(data)
	if err != nil {
		t.Errorf("delete webhook fail: %v", err)
		t.FailNow()
	}
	t.Logf(" create : %v, update: %v, delete : %v", res.Create, res.Update, res.Delete)
	for _, webhook := range res.Create {
		t.Logf(" webhook id: %v, webhook Name : %v", webhook.ID, webhook.Name)
	}

	for _, webhook := range res.Update {
		t.Logf(" order id: %v, webhook name : %v", webhook.ID, webhook.Name)
	}

	for _, webhook := range res.Delete {
		t.Logf(" webhook id: %v, webhook status : %v", webhook.ID, webhook.Status)
	}
}
