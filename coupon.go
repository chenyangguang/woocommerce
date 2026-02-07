package woocommerce

import (
	"fmt"
)

const (
	couponsBasePath = "coupons"
)

type CouponService interface {
	Create(coupon Coupon) (*Coupon, error)
	Get(couponID int64, options interface{}) (*Coupon, error)
	List(options interface{}) ([]Coupon, error)
	Update(coupon *Coupon) (*Coupon, error)
	Delete(couponID int64, options interface{}) (*Coupon, error)
	Batch(data CouponBatchOption) (*CouponBatchResource, error)
}

type Coupon struct {
	ID                int64                   `json:"id,omitempty"`
	Code              string                  `json:"code,omitempty"`
	Amount            string                  `json:"amount,omitempty"`
	DateCreated       string                  `json:"date_created,omitempty"`
	DateCreatedGmt    string                  `json:"date_created_gmt,omitempty"`
	DateModified      string                  `json:"date_modified,omitempty"`
	DateModifiedGmt   string                  `json:"date_modified_gmt,omitempty"`
	DiscountType      string                  `json:"discount_type,omitempty"`
	Description       string                  `json:"description,omitempty"`
	ExcludeSaleItems  bool                    `json:"exclude_sale_items,omitempty"`
	ExpiryDate        string                  `json:"expiry_date,omitempty"`
	FreeShipping      bool                    `json:"free_shipping,omitempty"`
	IndividualUse     bool                    `json:"individual_use,omitempty"`
	Length            int                     `json:"length,omitempty"`
	Limit             int64                   `json:"limit,omitempty"`
	MinimumAmount     string                  `json:"minimum_amount,omitempty"`
	MaximumAmount     string                  `json:"maximum_amount,omitempty"`
	UsageCount        int64                   `json:"usage_count,omitempty"`
	UsedBy            []string                `json:"used_by,omitempty"`
	EmailRestrictions CouponEmailRestrictions `json:"email_restrictions,omitempty"`
}

type CouponEmailRestrictions struct {
	Emails []string `json:"emails,omitempty"`
}

// CouponListOption list all the coupon list option request params
type CouponListOption struct {
	ListOptions
	Search  string  `url:"search,omitempty"`
	After   string  `url:"after,omitempty"`
	Before  string  `url:"before,omitempty"`
	Exclude []int64 `url:"exclude,omitempty"`
	Include []int64 `url:"include,omitempty"`
	Offset  int     `url:"offset,omitempty"`
	Order   string  `url:"order,omitempty"`
	Orderby string  `url:"orderby,omitempty"`
	Code    string  `url:"code,omitempty"`
}

type CouponBatchOption struct {
	Create []Coupon `json:"create,omitempty"`
	Update []Coupon `json:"update,omitempty"`
	Delete []int64  `json:"delete,omitempty"`
}

type CouponBatchResource struct {
	Create []*Coupon `json:"create,omitempty"`
	Update []*Coupon `json:"update,omitempty"`
	Delete []*Coupon `json:"delete,omitempty"`
}

type CouponServiceOp struct {
	client *Client
}

func (c *CouponServiceOp) Create(coupon Coupon) (*Coupon, error) {
	path := fmt.Sprintf("%s", couponsBasePath)
	resource := new(Coupon)
	err := c.client.Post(path, coupon, &resource)
	return resource, err
}

// Get individual coupon
func (c *CouponServiceOp) Get(couponID int64, options interface{}) (*Coupon, error) {
	path := fmt.Sprintf("%s/%d", couponsBasePath, couponID)
	resource := new(Coupon)
	err := c.client.Get(path, resource, options)
	return resource, err
}

func (c *CouponServiceOp) List(options interface{}) ([]Coupon, error) {
	path := fmt.Sprintf("%s", couponsBasePath)
	resource := make([]Coupon, 0)
	err := c.client.Get(path, &resource, options)
	return resource, err
}

func (c *CouponServiceOp) Update(coupon *Coupon) (*Coupon, error) {
	path := fmt.Sprintf("%s/%d", couponsBasePath, coupon.ID)
	resource := new(Coupon)
	err := c.client.Put(path, coupon, &resource)
	return resource, err
}

func (c *CouponServiceOp) Delete(couponID int64, options interface{}) (*Coupon, error) {
	path := fmt.Sprintf("%s/%d", couponsBasePath, couponID)
	resource := new(Coupon)
	err := c.client.Delete(path, options, &resource)
	return resource, err
}

func (c *CouponServiceOp) Batch(data CouponBatchOption) (*CouponBatchResource, error) {
	path := fmt.Sprintf("%s/batch", couponsBasePath)
	resource := new(CouponBatchResource)
	err := c.client.Post(path, data, &resource)
	return resource, err
}
