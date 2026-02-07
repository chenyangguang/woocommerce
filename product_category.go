package woocommerce

import (
	"fmt"
)

const (
	productCategoriesBasePath = "products/categories"
)

type ProductCategoryService interface {
	Create(category ProductCategory) (*ProductCategory, error)
	Get(categoryID int64, options interface{}) (*ProductCategory, error)
	List(options interface{}) ([]ProductCategory, error)
	Update(category *ProductCategory) (*ProductCategory, error)
	Delete(categoryID int64, options interface{}) (*ProductCategory, error)
	Batch(data ProductCategoryBatchOption) (*ProductCategoryBatchResource, error)
}

type ProductCategory struct {
	ID          int64        `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Slug        string       `json:"slug,omitempty"`
	ParentID    int64        `json:"parent,omitempty"`
	Description string       `json:"description,omitempty"`
	Display     string       `json:"display,omitempty"`
	Image       ProductImage `json:"image,omitempty"`
	MenuOrder   int          `json:"menu_order,omitempty"`
	Count       int64        `json:"count,omitempty"`
	Links       Links        `json:"_links,omitempty"`
}

type ProductCategoryListOption struct {
	ListOptions
	Search    string  `url:"search,omitempty"`
	Exclude   []int64 `url:"exclude,omitempty"`
	Include   []int64 `url:"include,omitempty"`
	Offset    int     `url:"offset,omitempty"`
	Order     string  `url:"order,omitempty"`
	Orderby   string  `url:"orderby,omitempty"`
	HideEmpty bool    `url:"hide_empty,omitempty"`
	Parent    []int64 `url:"parent,omitempty"`
	Product   int64   `url:"product,omitempty"`
	Slug      string  `url:"slug,omitempty"`
}

type ProductCategoryBatchOption struct {
	Create []ProductCategory `json:"create,omitempty"`
	Update []ProductCategory `json:"update,omitempty"`
	Delete []int64           `json:"delete,omitempty"`
}

type ProductCategoryBatchResource struct {
	Create []*ProductCategory `json:"create,omitempty"`
	Update []*ProductCategory `json:"update,omitempty"`
	Delete []*ProductCategory `json:"delete,omitempty"`
}

type ProductCategoryServiceOp struct {
	client *Client
}

func (c *ProductCategoryServiceOp) List(options interface{}) ([]ProductCategory, error) {
	path := fmt.Sprintf("%s", productCategoriesBasePath)
	resource := make([]ProductCategory, 0)
	err := c.client.Get(path, &resource, options)
	return resource, err
}

func (c *ProductCategoryServiceOp) Create(category ProductCategory) (*ProductCategory, error) {
	path := fmt.Sprintf("%s", productCategoriesBasePath)
	resource := new(ProductCategory)
	err := c.client.Post(path, category, &resource)
	return resource, err
}

func (c *ProductCategoryServiceOp) Get(categoryID int64, options interface{}) (*ProductCategory, error) {
	path := fmt.Sprintf("%s/%d", productCategoriesBasePath, categoryID)
	resource := new(ProductCategory)
	err := c.client.Get(path, resource, options)
	return resource, err
}

func (c *ProductCategoryServiceOp) Update(category *ProductCategory) (*ProductCategory, error) {
	path := fmt.Sprintf("%s/%d", productCategoriesBasePath, category.ID)
	resource := new(ProductCategory)
	err := c.client.Put(path, category, &resource)
	return resource, err
}

func (c *ProductCategoryServiceOp) Delete(categoryID int64, options interface{}) (*ProductCategory, error) {
	path := fmt.Sprintf("%s/%d", productCategoriesBasePath, categoryID)
	resource := new(ProductCategory)
	err := c.client.Delete(path, options, &resource)
	return resource, err
}

func (c *ProductCategoryServiceOp) Batch(data ProductCategoryBatchOption) (*ProductCategoryBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productCategoriesBasePath)
	resource := new(ProductCategoryBatchResource)
	err := c.client.Post(path, data, &resource)
	return resource, err
}
