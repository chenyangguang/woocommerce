package woocommerce

import (
	"fmt"
)

const (
	customersBasePath = "customers"
)

type CustomerService interface {
	Create(customer Customer) (*Customer, error)
	Get(customerID int64, options interface{}) (*Customer, error)
	List(options interface{}) ([]Customer, error)
	Update(customer *Customer) (*Customer, error)
	Delete(customerID int64, options interface{}) (*Customer, error)
	Batch(data CustomerBatchOption) (*CustomerBatchResource, error)
	GetDownloads(customerID int64, options interface{}) ([]CustomerDownload, error)
}

type Customer struct {
	ID              int64           `json:"id,omitempty"`
	DateCreated     string          `json:"date_created,omitempty"`
	DateCreatedGmt  string          `json:"date_created_gmt,omitempty"`
	DateModified    string          `json:"date_modified,omitempty"`
	DateModifiedGmt string          `json:"date_modified_gmt,omitempty"`
	Email           string          `json:"email,omitempty"`
	FirstName       string          `json:"first_name,omitempty"`
	LastName        string          `json:"last_name,omitempty"`
	Username        string          `json:"username,omitempty"`
	Role            string          `json:"role,omitempty"`
	Password        string          `json:"password,omitempty"`
	Billing         CustomerAddress `json:"billing,omitempty"`
	Shipping        CustomerAddress `json:"shipping,omitempty"`
	AvatarUrl       string          `json:"avatar_url,omitempty"`
	LinkToWoo       string          `json:"link_to_woo,omitempty"`
	LinkToWP        string          `json:"link_to_wp,omitempty"`
	LinkToAstra     string          `json:"link_to_astra,omitempty"`
}

// CustomerListOption list all the customer list option request params
type CustomerListOption struct {
	ListOptions
	Search  string  `url:"search,omitempty"`
	Role    string  `url:"role,omitempty"`
	Exclude []int64 `url:"exclude,omitempty"`
	Include []int64 `url:"include,omitempty"`
	Offset  int     `url:"offset,omitempty"`
	Order   string  `url:"order,omitempty"`
	Orderby string  `url:"orderby,omitempty"`
}

type CustomerBatchOption struct {
	Create []Customer `json:"create,omitempty"`
	Update []Customer `json:"update,omitempty"`
	Delete []int64    `json:"delete,omitempty"`
}

type CustomerBatchResource struct {
	Create []*Customer `json:"create,omitempty"`
	Update []*Customer `json:"update,omitempty"`
	Delete []*Customer `json:"delete,omitempty"`
}

type CustomerAddress struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address_1,omitempty"`
	Address2  string `json:"address_2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type CustomerDownload struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	File          string `json:"file,omitempty"`
	FilePath      string `json:"file_path,omitempty"`
	DownloadLimit int64  `json:"download_limit,omitempty"`
	DownloadCount int64  `json:"downloads_remaining,omitempty"`
	Expiry        string `json:"access_expires,omitempty"`
	DateCreated   string `json:"date_created,omitempty"`
}

type CustomerServiceOp struct {
	client *Client
}

func (c *CustomerServiceOp) Create(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s", customersBasePath)
	resource := new(Customer)
	err := c.client.Post(path, customer, &resource)
	return resource, err
}

func (c *CustomerServiceOp) List(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s", customersBasePath)
	resource := make([]Customer, 0)
	err := c.client.Get(path, &resource, options)
	return resource, err
}

// Get individual customer
func (c *CustomerServiceOp) Get(customerID int64, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%d", customersBasePath, customerID)
	resource := new(Customer)
	err := c.client.Get(path, resource, options)
	return resource, err
}

func (c *CustomerServiceOp) Update(customer *Customer) (*Customer, error) {
	path := fmt.Sprintf("%s/%d", customersBasePath, customer.ID)
	resource := new(Customer)
	err := c.client.Put(path, customer, &resource)
	return resource, err
}

func (c *CustomerServiceOp) Delete(customerID int64, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%d", customersBasePath, customerID)
	resource := new(Customer)
	err := c.client.Delete(path, options, &resource)
	return resource, err
}

func (c *CustomerServiceOp) Batch(data CustomerBatchOption) (*CustomerBatchResource, error) {
	path := fmt.Sprintf("%s/batch", customersBasePath)
	resource := new(CustomerBatchResource)
	err := c.client.Post(path, data, &resource)
	return resource, err
}

// GetDownloads retrieves downloads for a customer
func (c *CustomerServiceOp) GetDownloads(customerID int64, options interface{}) ([]CustomerDownload, error) {
	path := fmt.Sprintf("%s/%d/downloads", customersBasePath, customerID)
	resource := make([]CustomerDownload, 0)
	err := c.client.Get(path, &resource, options)
	return resource, err
}
