package woocommerce

import (
	"fmt"
)

const (
	productTagsBasePath = "products/tags"
)

type ProductTagService interface {
	Create(tag ProductTag) (*ProductTag, error)
	Get(tagID int64, options interface{}) (*ProductTag, error)
	List(options interface{}) ([]ProductTag, error)
	Update(tag *ProductTag) (*ProductTag, error)
	Delete(tagID int64, options interface{}) (*ProductTag, error)
	Batch(data ProductTagBatchOption) (*ProductTagBatchResource, error)
}

type ProductTag struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Count       int64  `json:"count,omitempty"`
	Links       Links  `json:"_links,omitempty"`
}

type ProductTagListOption struct {
	ListOptions
	HideEmpty bool   `url:"hide_empty,omitempty"`
	Order     string `url:"order,omitempty"`
	Orderby   string `url:"orderby,omitempty"`
}

type ProductTagBatchOption struct {
	Create []ProductTag `json:"create,omitempty"`
	Update []ProductTag `json:"update,omitempty"`
	Delete []int64      `json:"delete,omitempty"`
}

type ProductTagBatchResource struct {
	Create []*ProductTag `json:"create,omitempty"`
	Update []*ProductTag `json:"update,omitempty"`
	Delete []*ProductTag `json:"delete,omitempty"`
}

type ProductTagServiceOp struct {
	client *Client
}

func (t *ProductTagServiceOp) List(options interface{}) ([]ProductTag, error) {
	path := fmt.Sprintf("%s", productTagsBasePath)
	resource := make([]ProductTag, 0)
	err := t.client.Get(path, &resource, options)
	return resource, err
}

func (t *ProductTagServiceOp) Create(tag ProductTag) (*ProductTag, error) {
	path := fmt.Sprintf("%s", productTagsBasePath)
	resource := new(ProductTag)
	err := t.client.Post(path, tag, &resource)
	return resource, err
}

func (t *ProductTagServiceOp) Get(tagID int64, options interface{}) (*ProductTag, error) {
	path := fmt.Sprintf("%s/%d", productTagsBasePath, tagID)
	resource := new(ProductTag)
	err := t.client.Get(path, resource, options)
	return resource, err
}

func (t *ProductTagServiceOp) Update(tag *ProductTag) (*ProductTag, error) {
	path := fmt.Sprintf("%s/%d", productTagsBasePath, tag.ID)
	resource := new(ProductTag)
	err := t.client.Put(path, tag, &resource)
	return resource, err
}

func (t *ProductTagServiceOp) Delete(tagID int64, options interface{}) (*ProductTag, error) {
	path := fmt.Sprintf("%s/%d", productTagsBasePath, tagID)
	resource := new(ProductTag)
	err := t.client.Delete(path, options, &resource)
	return resource, err
}

func (t *ProductTagServiceOp) Batch(data ProductTagBatchOption) (*ProductTagBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productTagsBasePath)
	resource := new(ProductTagBatchResource)
	err := t.client.Post(path, data, &resource)
	return resource, err
}
