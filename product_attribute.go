package woocommerce

import (
	"fmt"
)

const (
	productAttributesBasePath = "products/attributes"
)

type ProductAttributeService interface {
	Create(attribute ProductAttributeData) (*ProductAttributeData, error)
	Get(attributeID int64, options interface{}) (*ProductAttributeData, error)
	List(options interface{}) ([]ProductAttributeData, error)
	Update(attribute *ProductAttributeData) (*ProductAttributeData, error)
	Delete(attributeID int64, options interface{}) (*ProductAttributeData, error)
	Batch(data ProductAttributeBatchOption) (*ProductAttributeBatchResource, error)
}

type ProductAttributeData struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Type        string `json:"type,omitempty"`
	Order       int    `json:"order,omitempty"`
	HasArchives bool   `json:"has_archives,omitempty"`
	Visible     bool   `json:"visible,omitempty"`
}

type ProductAttributeListOption struct {
	ListOptions
	Hidden bool `url:"hidden,omitempty"`
}

type ProductAttributeBatchOption struct {
	Create []ProductAttributeData `json:"create,omitempty"`
	Update []ProductAttributeData `json:"update,omitempty"`
	Delete []int64                `json:"delete,omitempty"`
}

type ProductAttributeBatchResource struct {
	Create []*ProductAttributeData `json:"create,omitempty"`
	Update []*ProductAttributeData `json:"update,omitempty"`
	Delete []*ProductAttributeData `json:"delete,omitempty"`
}

type ProductAttributeServiceOp struct {
	client *Client
}

func (a *ProductAttributeServiceOp) List(options interface{}) ([]ProductAttributeData, error) {
	path := fmt.Sprintf("%s", productAttributesBasePath)
	resource := make([]ProductAttributeData, 0)
	err := a.client.Get(path, &resource, options)
	return resource, err
}

func (a *ProductAttributeServiceOp) Create(attribute ProductAttributeData) (*ProductAttributeData, error) {
	path := fmt.Sprintf("%s", productAttributesBasePath)
	resource := new(ProductAttributeData)
	err := a.client.Post(path, attribute, &resource)
	return resource, err
}

func (a *ProductAttributeServiceOp) Get(attributeID int64, options interface{}) (*ProductAttributeData, error) {
	path := fmt.Sprintf("%s/%d", productAttributesBasePath, attributeID)
	resource := new(ProductAttributeData)
	err := a.client.Get(path, resource, options)
	return resource, err
}

func (a *ProductAttributeServiceOp) Update(attribute *ProductAttributeData) (*ProductAttributeData, error) {
	path := fmt.Sprintf("%s/%d", productAttributesBasePath, attribute.ID)
	resource := new(ProductAttributeData)
	err := a.client.Put(path, attribute, &resource)
	return resource, err
}

func (a *ProductAttributeServiceOp) Delete(attributeID int64, options interface{}) (*ProductAttributeData, error) {
	path := fmt.Sprintf("%s/%d", productAttributesBasePath, attributeID)
	resource := new(ProductAttributeData)
	err := a.client.Delete(path, options, &resource)
	return resource, err
}

func (a *ProductAttributeServiceOp) Batch(data ProductAttributeBatchOption) (*ProductAttributeBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productAttributesBasePath)
	resource := new(ProductAttributeBatchResource)
	err := a.client.Post(path, data, &resource)
	return resource, err
}
