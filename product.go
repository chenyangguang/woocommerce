package gowooco

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(prev|next|first|last)" *$`)

// ProductService allows you to create, view, update, and delete individual, or a batch, of products
// https://woocommerce.github.io/woocommerce-rest-api-docs/#products
type ProductService interface {
	Create()
	Get()
	Delete()
	List()
	Update()
	BatchUpdate()
}

// Product represent WooCommerce Product
// https://woocommerce.github.io/woocommerce-rest-api-docs/#product-properties
type Product struct {
}

type ProductServiceOp struct {
	client *Client
}

// Pagination of results
type Pagination struct {
	NextPageOptions     *ListOptions
	PreviousPageOptions *ListOptions
	FirstPageOptions    *ListOptions
	LastPageOptions     *ListOptions
}

func (p *ProductServiceOp) Create() {

}

func (p *ProductServiceOp) Get() {

}

func (p *ProductServiceOp) Delete() {

}

func (p *ProductServiceOp) List() {

}

func (p *ProductServiceOp) Update() {

}

func (p *ProductServiceOp) BatchUpdate() {

}

// extractPagination extracts pagination info from linkHeader.
// Details on the format are here:
// https://woocommerce.github.io/woocommerce-rest-api-docs/#pagination
// Link: <https://www.example.com/wp-json/wc/v3/products?page=2>; rel="next",
// <https://www.example.com/wp-json/wc/v3/products?page=3>; rel="last"`
func extractPagination(linkHeader string) (*Pagination, error) {
	pagination := new(Pagination)

	if linkHeader == "" {
		return pagination, nil
	}

	for _, link := range strings.Split(linkHeader, ",") {
		match := linkRegex.FindStringSubmatch(link)
		// Make sure the link is not empty or invalid
		println("mm", len(match))
		if len(match) != 4 {
			// We expect 3 values:
			// match[0] = full match
			// match[1] is the URL and match[2] is either 'previous' or 'next', 'first', 'last'
			err := ResponseDecodingError{
				Message: "could not extract pagination link header",
			}
			return nil, err
		}

		rel, err := url.Parse(match[1])
		if err != nil {
			err = ResponseDecodingError{
				Message: "pagination does not contain a valid URL",
			}
			return nil, err
		}

		params, err := url.ParseQuery(rel.RawQuery)
		if err != nil {
			return nil, err
		}

		paginationListOptions := ListOptions{}

		page := params.Get("page")
		if page != "" {
			paginationListOptions.Page, err = strconv.Atoi(params.Get("page"))
			if err != nil {
				return nil, err
			}
		}

		switch match[2] {
		case "next":
			pagination.NextPageOptions = &paginationListOptions
		case "prev":
			pagination.PreviousPageOptions = &paginationListOptions
		case "first":
			pagination.FirstPageOptions = &paginationListOptions
		case "last":
			pagination.LastPageOptions = &paginationListOptions
		}

	}

	return pagination, nil
}
