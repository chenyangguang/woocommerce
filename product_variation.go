package woocommerce

import (
	"fmt"
)

const (
	productVariationsBasePath = "products/%d/variations"
)

// ProductVariationService allows you to create, view, update, and delete individual, or a batch, of product variations
type ProductVariationService interface {
	Create(productID int64, variation ProductVariation) (*ProductVariation, error)
	Get(productID int64, variationID int64, options interface{}) (*ProductVariation, error)
	List(productID int64, options interface{}) ([]ProductVariation, error)
	Update(productID int64, variation *ProductVariation) (*ProductVariation, error)
	Delete(productID int64, variationID int64, options interface{}) (*ProductVariation, error)
	Batch(productID int64, data ProductVariationBatchOption) (*ProductVariationBatchResource, error)
}

type ProductVariation struct {
	ID               int64              `json:"id,omitempty"`
	DateCreated      string             `json:"date_created,omitempty"`
	DateCreatedGmt   string             `json:"date_created_gmt,omitempty"`
	DateModified     string             `json:"date_modified,omitempty"`
	DateModifiedGmt  string             `json:"date_modified_gmt,omitempty"`
	Permalink        string             `json:"permalink,omitempty"`
	SKU              string             `json:"sku,omitempty"`
	Price            string             `json:"price,omitempty"`
	RegularPrice     string             `json:"regular_price,omitempty"`
	SalePrice        string             `json:"sale_price,omitempty"`
	DateOnSaleFrom   string             `json:"date_on_sale_from,omitempty"`
	DateOnSaleTo     string             `json:"date_on_sale_to,omitempty"`
	Status           string             `json:"status,omitempty"`
	Virtual          bool               `json:"virtual,omitempty"`
	Downloadable     bool               `json:"downloadable,omitempty"`
	ManageStock      bool               `json:"manage_stock,omitempty"`
	StockQuantity    string             `json:"stock_quantity,omitempty"`
	StockStatus      string             `json:"stock_status,omitempty"`
	Backorders       string             `json:"backorders,omitempty"`
	LowStockAmount   string             `json:"low_stock_amount,omitempty"`
	SoldIndividually bool               `json:"sold_individually,omitempty"`
	Weight           string             `json:"weight,omitempty"`
	Length           string             `json:"length,omitempty"`
	Width            string             `json:"width,omitempty"`
	Height           string             `json:"height,omitempty"`
	Dimensions       map[string]string  `json:"dimensions,omitempty"`
	ShippingClass    string             `json:"shipping_class,omitempty"`
	ShippingClassID  int64              `json:"shipping_class_id,omitempty"`
	Image            ProductImage       `json:"image,omitempty"`
	Attributes       []ProductAttribute `json:"attributes,omitempty"`
	MetaData         []MetaData         `json:"meta_data,omitempty"`
	MenuOrder        int                `json:"menu_order,omitempty"`
	Links            Links              `json:"_links,omitempty"`
}

type ProductVariationListOption struct {
	ListOptions
}

type ProductVariationBatchOption struct {
	Create []ProductVariation `json:"create,omitempty"`
	Update []ProductVariation `json:"update,omitempty"`
	Delete []int64            `json:"delete,omitempty"`
}

type ProductVariationBatchResource struct {
	Create []*ProductVariation `json:"create,omitempty"`
	Update []*ProductVariation `json:"update,omitempty"`
	Delete []*ProductVariation `json:"delete,omitempty"`
}

type ProductVariationServiceOp struct {
	client *Client
}

func (p *ProductVariationServiceOp) List(productID int64, options interface{}) ([]ProductVariation, error) {
	path := fmt.Sprintf(productVariationsBasePath, productID)
	resource := make([]ProductVariation, 0)
	err := p.client.Get(path, &resource, options)
	return resource, err
}

func (p *ProductVariationServiceOp) Create(productID int64, variation ProductVariation) (*ProductVariation, error) {
	path := fmt.Sprintf(productVariationsBasePath, productID)
	resource := new(ProductVariation)
	err := p.client.Post(path, variation, &resource)
	return resource, err
}

func (p *ProductVariationServiceOp) Get(productID int64, variationID int64, options interface{}) (*ProductVariation, error) {
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(productVariationsBasePath, productID), variationID)
	resource := new(ProductVariation)
	err := p.client.Get(path, resource, options)
	return resource, err
}

func (p *ProductVariationServiceOp) Update(productID int64, variation *ProductVariation) (*ProductVariation, error) {
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(productVariationsBasePath, productID), variation.ID)
	resource := new(ProductVariation)
	err := p.client.Put(path, variation, &resource)
	return resource, err
}

func (p *ProductVariationServiceOp) Delete(productID int64, variationID int64, options interface{}) (*ProductVariation, error) {
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(productVariationsBasePath, productID), variationID)
	resource := new(ProductVariation)
	err := p.client.Delete(path, options, &resource)
	return resource, err
}

func (p *ProductVariationServiceOp) Batch(productID int64, data ProductVariationBatchOption) (*ProductVariationBatchResource, error) {
	path := fmt.Sprintf("%s/batch", fmt.Sprintf(productVariationsBasePath, productID))
	resource := new(ProductVariationBatchResource)
	err := p.client.Post(path, data, &resource)
	return resource, err
}
