package woocommerce

import (
	"fmt"
)

const (
	productShippingClassesBasePath = "products/shipping_classes"
)

type ProductShippingClassService interface {
	Create(shippingClass ProductShippingClass) (*ProductShippingClass, error)
	Get(shippingClassID int64, options interface{}) (*ProductShippingClass, error)
	List(options interface{}) ([]ProductShippingClass, error)
	Update(shippingClass *ProductShippingClass) (*ProductShippingClass, error)
	Delete(shippingClassID int64, options interface{}) (*ProductShippingClass, error)
	Batch(data ProductShippingClassBatchOption) (*ProductShippingClassBatchResource, error)
}

type ProductShippingClass struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Slug  string `json:"slug,omitempty"`
	Count int64  `json:"count,omitempty"`
	Links Links  `json:"_links,omitempty"`
}

type ProductShippingClassListOption struct {
	ListOptions
	HideEmpty bool   `url:"hide_empty,omitempty"`
	Order     string `url:"order,omitempty"`
	Orderby   string `url:"orderby,omitempty"`
}

type ProductShippingClassBatchOption struct {
	Create []ProductShippingClass `json:"create,omitempty"`
	Update []ProductShippingClass `json:"update,omitempty"`
	Delete []int64                `json:"delete,omitempty"`
}

type ProductShippingClassBatchResource struct {
	Create []*ProductShippingClass `json:"create,omitempty"`
	Update []*ProductShippingClass `json:"update,omitempty"`
	Delete []*ProductShippingClass `json:"delete,omitempty"`
}

type ProductShippingClassServiceOp struct {
	client *Client
}

func (s *ProductShippingClassServiceOp) List(options interface{}) ([]ProductShippingClass, error) {
	path := fmt.Sprintf("%s", productShippingClassesBasePath)
	resource := make([]ProductShippingClass, 0)
	err := s.client.Get(path, &resource, options)
	return resource, err
}

func (s *ProductShippingClassServiceOp) Create(shippingClass ProductShippingClass) (*ProductShippingClass, error) {
	path := fmt.Sprintf("%s", productShippingClassesBasePath)
	resource := new(ProductShippingClass)
	err := s.client.Post(path, shippingClass, &resource)
	return resource, err
}

func (s *ProductShippingClassServiceOp) Get(shippingClassID int64, options interface{}) (*ProductShippingClass, error) {
	path := fmt.Sprintf("%s/%d", productShippingClassesBasePath, shippingClassID)
	resource := new(ProductShippingClass)
	err := s.client.Get(path, resource, options)
	return resource, err
}

func (s *ProductShippingClassServiceOp) Update(shippingClass *ProductShippingClass) (*ProductShippingClass, error) {
	path := fmt.Sprintf("%s/%d", productShippingClassesBasePath, shippingClass.ID)
	resource := new(ProductShippingClass)
	err := s.client.Put(path, shippingClass, &resource)
	return resource, err
}

func (s *ProductShippingClassServiceOp) Delete(shippingClassID int64, options interface{}) (*ProductShippingClass, error) {
	path := fmt.Sprintf("%s/%d", productShippingClassesBasePath, shippingClassID)
	resource := new(ProductShippingClass)
	err := s.client.Delete(path, options, &resource)
	return resource, err
}

func (s *ProductShippingClassServiceOp) Batch(data ProductShippingClassBatchOption) (*ProductShippingClassBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productShippingClassesBasePath)
	resource := new(ProductShippingClassBatchResource)
	err := s.client.Post(path, data, &resource)
	return resource, err
}
