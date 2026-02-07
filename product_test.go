package woocommerce

import (
	"testing"
	"time"
)

func TestProductServiceOp_List(t *testing.T) {
	options := ProductListOption{
		ListOptions: ListOptions{
			Context: "view",
			After:   "2021-01-01T00:00:00",
			Before:  "2023-12-31T23:59:59",
			Order:   "desc",
			Orderby: "date",
			Page:    1,
			PerPage: 10,
		},
		Type: "simple",
	}
	products, err := client.Product.List(options)
	if err != nil {
		t.Logf("error listing products: %v", err)
	}
	for _, product := range products {
		t.Logf("product: id=%d, name=%s, type=%s, price=%s", product.ID, product.Name, product.Type, product.Price)
	}
}

func TestProductServiceOp_Create(t *testing.T) {
	product := Product{
		Name:             "Test Product " + time.Now().Format("20060102150405"),
		Type:             "simple",
		RegularPrice:     "29.99",
		Description:      "A test product",
		ShortDescription: "Short test product description",
		SKU:              "test-sku-" + time.Now().Format("20060102150405"),
		ManageStock:      true,
		StockQuantity:    "100",
		Status:           "publish",
	}
	res, err := client.Product.Create(product)
	if err != nil {
		t.Logf("create product error: %v", err)
	} else {
		t.Logf("created product: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductServiceOp_Get(t *testing.T) {
	product, err := client.Product.Get(1, nil)
	if err != nil {
		t.Logf("get product error: %v", err)
	} else {
		t.Logf("got product: id=%d, name=%s, price=%s", product.ID, product.Name, product.Price)
	}
}

func TestProductServiceOp_Update(t *testing.T) {
	product, err := client.Product.Get(1, nil)
	if product == nil || err != nil {
		t.Skip("skipping update test: cannot get product 1")
		return
	}
	product.Description = "Updated description " + time.Now().Format("20060102150405")
	res, err := client.Product.Update(product)
	if err != nil {
		t.Logf("update product error: %v", err)
	} else {
		t.Logf("updated product: id=%d, description=%s", res.ID, res.Description)
	}
}

func TestProductServiceOp_Delete(t *testing.T) {
	product := Product{
		Name:         "Test Product to Delete " + time.Now().Format("20060102150405"),
		Type:         "simple",
		RegularPrice: "19.99",
		Status:       "publish",
	}
	created, err := client.Product.Create(product)
	if err != nil {
		t.Skipf("skipping delete test: cannot create product: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.Product.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete product error: %v", err)
	} else {
		t.Logf("deleted product: id=%d", res.ID)
	}
}

func TestProductServiceOp_Batch(t *testing.T) {
	timeNow := time.Now().Format("20060102150405")
	data := ProductBatchOption{
		Create: []Product{
			{
				Name:         "Batch Product 1 " + timeNow,
				Type:         "simple",
				RegularPrice: "10.99",
				Status:       "publish",
			},
			{
				Name:         "Batch Product 2 " + timeNow,
				Type:         "simple",
				RegularPrice: "20.99",
				Status:       "publish",
			},
		},
	}
	res, err := client.Product.Batch(data)
	if err != nil {
		t.Logf("batch products error: %v", err)
	} else {
		t.Logf("batch created %d products", len(res.Create))
		for _, p := range res.Create {
			t.Logf("batch created product: id=%d, name=%s", p.ID, p.Name)
		}
	}
}
