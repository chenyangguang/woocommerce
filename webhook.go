package woocommerce

import (
	"fmt"
)

const (
	webhooksBasePath = "webhooks"
)

// WebhookService is an interface for interfacing with the webhook endpoints of
// the WooCommerce webhooks restful API
// https://woocommerce.github.io/woocommerce-rest-api-docs/#webhooks
type WebhookService interface {
	List(options interface{}) ([]Webhook, error)
	Create(webhook Webhook) (*Webhook, error)
	Get(webhookID int64, options interface{}) (*Webhook, error)
	Update(webhook *Webhook) (*Webhook, error)
	Delete(webhookID int64, options interface{}) (*Webhook, error)
	Batch(data WebhookBatchOption) (*WebhookBatchResource, error)
}

// WebhookServiceOp handles communication with the webhooks related methods of WooCommerce restful api
type WebhookServiceOp struct {
	client *Client
}

// Webhook represent a  wooCommerce webhook's All  properties columns
type Webhook struct {
	ID              int64    `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Status          string   `json:"status,omitempty"`
	Topic           string   `json:"topic,omitempty"`
	Resource        string   `json:"resource,omitempty"`
	Event           string   `json:"event,omitempty"`
	Hooks           []string `json:"hooks,omitempty"`
	DeliveryUrl     string   `json:"delivery_url,omitempty"`
	Secret          string   `json:"secret,omitempty"`
	DateCreated     string   `json:"date_created,omitempty"`
	DateCreatedGmt  string   `json:"date_created_gmt,omitempty"`
	DateModified    string   `json:"date_modified,omitempty"`
	DateModifiedGmt string   `json:"date_modified_gmt,omitempty"`
	Links           Links    `json:"_links,omitempty"`
}

// WebhookListOption config webhook's List method request option
type WebhookListOption struct {
	ListOptions
	Status string `json:"status,omitempty"`
}

// WebhookDeleteOption config webhook's Delete operation option
type WebhookDeleteOption struct {
	Force bool
}

// OrderBatchOption setting  operate for order in batch way
type WebhookBatchOption struct {
	Create []Webhook `json:"create,omitempty"`
	Update []Webhook `json:"update,omitempty"`
	Delete []int64   `json:"delete,omitempty"`
}

// WebhookBatchResource conservation the response struct for WebhookBatchOption's request
type WebhookBatchResource struct {
	Create []*Webhook `json:"create,omitempty"`
	Update []*Webhook `json:"update,omitempty"`
	Delete []*Webhook `json:"delete,omitempty"`
}

// List return multiple webhooks
// https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-webhooks
func (w *WebhookServiceOp) List(options interface{}) ([]Webhook, error) {
	path := fmt.Sprintf("%s", webhooksBasePath)
	resource := make([]Webhook, 0)
	err := w.client.Get(path, &resource, options)
	return resource, err
}

// Create handle create a new webhook.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#create-a-webhook
func (w *WebhookServiceOp) Create(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s", webhooksBasePath)
	resource := new(Webhook)
	err := w.client.Post(path, webhook, &resource)
	return resource, err
}

// Get implement for retrieve and view a specific webhook
// https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-webhook
func (w *WebhookServiceOp) Get(webhookID int64, options interface{}) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d", webhooksBasePath, webhookID)
	resource := new(Webhook)
	err := w.client.Get(path, &resource, options)
	return resource, err
}

// Update method allow you to make changes to a webhook
// https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-webhook
func (w *WebhookServiceOp) Update(webhook *Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d", webhooksBasePath, webhook.ID)
	resource := new(Webhook)
	err := w.client.Put(path, webhook, &resource)

	return resource, err
}

// Delete delete a webhook
// https://woocommerce.github.io/woocommerce-rest-api-docs/#delete-a-webhook
func (w *WebhookServiceOp) Delete(webhookID int64, options interface{}) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d", webhooksBasePath, webhookID)
	resource := new(Webhook)
	err := w.client.Delete(path, options, &resource)
	return resource, err
}

// Batch helps you to batch create, update and delete multiple webhooks
// WooCommerce docs Notes : By default it's limited to up to 100 objects to be created, updated or deleted.
// reference :
// https://woocommerce.github.io/woocommerce-rest-api-docs/#batch-update-webhooks
func (w *WebhookServiceOp) Batch(data WebhookBatchOption) (*WebhookBatchResource, error) {
	path := fmt.Sprintf("%s/batch", webhooksBasePath)
	resource := new(WebhookBatchResource)
	err := w.client.Post(path, data, &resource)
	return resource, err
}
