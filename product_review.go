package woocommerce

import (
	"fmt"
)

const (
	productReviewsBasePath = "products/reviews"
)

type ProductReviewService interface {
	Create(review ProductReview) (*ProductReview, error)
	Get(reviewID int64, options interface{}) (*ProductReview, error)
	List(options interface{}) ([]ProductReview, error)
	Update(review *ProductReview) (*ProductReview, error)
	Delete(reviewID int64, options interface{}) (*ProductReview, error)
	Batch(data ProductReviewBatchOption) (*ProductReviewBatchResource, error)
}

type ProductReviewListOption struct {
	ListOptions
	Search  string  `url:"search,omitempty"`
	Exclude []int64 `url:"exclude,omitempty"`
	Include []int64 `url:"include,omitempty"`
	Offset  int     `url:"offset,omitempty"`
	Order   string  `url:"order,omitempty"`
	Orderby string  `url:"orderby,omitempty"`
}

type ProductReviewBatchOption struct {
	Create []ProductReview `json:"create,omitempty"`
	Update []ProductReview `json:"update,omitempty"`
	Delete []int64         `json:"delete,omitempty"`
}

type ProductReviewBatchResource struct {
	Create []*ProductReview `json:"create,omitempty"`
	Update []*ProductReview `json:"update,omitempty"`
	Delete []*ProductReview `json:"delete,omitempty"`
}

type ProductReviewServiceOp struct {
	client *Client
}

func (r *ProductReviewServiceOp) List(options interface{}) ([]ProductReview, error) {
	path := fmt.Sprintf("%s", productReviewsBasePath)
	resource := make([]ProductReview, 0)
	err := r.client.Get(path, &resource, options)
	return resource, err
}

func (r *ProductReviewServiceOp) Create(review ProductReview) (*ProductReview, error) {
	path := fmt.Sprintf("%s", productReviewsBasePath)
	resource := new(ProductReview)
	err := r.client.Post(path, review, &resource)
	return resource, err
}

func (r *ProductReviewServiceOp) Get(reviewID int64, options interface{}) (*ProductReview, error) {
	path := fmt.Sprintf("%s/%d", productReviewsBasePath, reviewID)
	resource := new(ProductReview)
	err := r.client.Get(path, resource, options)
	return resource, err
}

func (r *ProductReviewServiceOp) Update(review *ProductReview) (*ProductReview, error) {
	path := fmt.Sprintf("%s/%d", productReviewsBasePath, review.ID)
	resource := new(ProductReview)
	err := r.client.Put(path, review, &resource)
	return resource, err
}

func (r *ProductReviewServiceOp) Delete(reviewID int64, options interface{}) (*ProductReview, error) {
	path := fmt.Sprintf("%s/%d", productReviewsBasePath, reviewID)
	resource := new(ProductReview)
	err := r.client.Delete(path, options, &resource)
	return resource, err
}

func (r *ProductReviewServiceOp) Batch(data ProductReviewBatchOption) (*ProductReviewBatchResource, error) {
	path := fmt.Sprintf("%s/batch", productReviewsBasePath)
	resource := new(ProductReviewBatchResource)
	err := r.client.Post(path, data, &resource)
	return resource, err
}
