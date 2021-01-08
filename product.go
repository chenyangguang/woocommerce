package gowooco

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
