package woocommerce

import (
	"testing"
	"time"
)

func TestProductVariationServiceOp_List(t *testing.T) {
	productID := int64(1)
	variations, err := client.ProductVariation.List(productID, nil)
	if err != nil {
		t.Logf("error listing variations: %v", err)
	}
	for _, variation := range variations {
		t.Logf("variation: id=%d, sku=%s, price=%s", variation.ID, variation.SKU, variation.Price)
	}
}

func TestProductVariationServiceOp_Create(t *testing.T) {
	productID := int64(1)
	variation := ProductVariation{
		SKU:           "var-test-" + time.Now().Format("20060102150405"),
		RegularPrice:  "15.99",
		SalePrice:     "12.99",
		ManageStock:   true,
		StockQuantity: "50",
		Status:        "publish",
	}
	res, err := client.ProductVariation.Create(productID, variation)
	if err != nil {
		t.Logf("create variation error: %v", err)
	} else {
		t.Logf("created variation: id=%d, sku=%s", res.ID, res.SKU)
	}
}

func TestProductVariationServiceOp_Get(t *testing.T) {
	productID := int64(1)
	variationID := int64(1)
	variation, err := client.ProductVariation.Get(productID, variationID, nil)
	if err != nil {
		t.Logf("get variation error: %v", err)
	} else {
		t.Logf("got variation: id=%d, sku=%s", variation.ID, variation.SKU)
	}
}

func TestProductVariationServiceOp_Update(t *testing.T) {
	productID := int64(1)
	variationID := int64(1)
	variation, err := client.ProductVariation.Get(productID, variationID, nil)
	if variation == nil || err != nil {
		t.Skip("skipping update test: cannot get variation")
		return
	}
	variation.RegularPrice = "25.99"
	res, err := client.ProductVariation.Update(productID, variation)
	if err != nil {
		t.Logf("update variation error: %v", err)
	} else {
		t.Logf("updated variation: id=%d, price=%s", res.ID, res.RegularPrice)
	}
}

func TestProductVariationServiceOp_Delete(t *testing.T) {
	productID := int64(1)
	variation := ProductVariation{
		SKU:          "delete-var-" + time.Now().Format("20060102150405"),
		RegularPrice: "9.99",
		Status:       "publish",
	}
	created, err := client.ProductVariation.Create(productID, variation)
	if err != nil {
		t.Skipf("skipping delete test: cannot create variation: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	_, err = client.ProductVariation.Delete(productID, created.ID, optionsDel)
	if err != nil {
		t.Logf("delete variation error: %v", err)
	} else {
		t.Logf("deleted variation: id=%d", created.ID)
	}
}

func TestProductVariationServiceOp_Batch(t *testing.T) {
	productID := int64(1)
	timeNow := time.Now().Format("20060102150405")
	data := ProductVariationBatchOption{
		Create: []ProductVariation{
			{
				SKU:          "batch-var1-" + timeNow,
				RegularPrice: "11.99",
				Status:       "publish",
			},
			{
				SKU:          "batch-var2-" + timeNow,
				RegularPrice: "22.99",
				Status:       "publish",
			},
		},
	}
	res, err := client.ProductVariation.Batch(productID, data)
	if err != nil {
		t.Logf("batch variations error: %v", err)
	} else {
		t.Logf("batch created %d variations", len(res.Create))
		for _, v := range res.Create {
			t.Logf("batch created variation: id=%d, sku=%s", v.ID, v.SKU)
		}
	}
}
