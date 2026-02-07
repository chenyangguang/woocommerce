package woocommerce

import (
	"testing"
)

func TestProductShippingClassServiceOp_List(t *testing.T) {
	shippingClasses, err := client.ProductShippingClass.List(nil)
	if err != nil {
		t.Logf("error listing shipping classes: %v", err)
	}
	for _, sc := range shippingClasses {
		t.Logf("shipping class: id=%d, name=%s, slug=%s", sc.ID, sc.Name, sc.Slug)
	}
}

func TestProductShippingClassServiceOp_Create(t *testing.T) {
	shippingClass := ProductShippingClass{
		Name: "Test Shipping Class",
		Slug: "test-shipping-class",
	}
	res, err := client.ProductShippingClass.Create(shippingClass)
	if err != nil {
		t.Logf("create shipping class error: %v", err)
	} else {
		t.Logf("created shipping class: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductShippingClassServiceOp_Get(t *testing.T) {
	shippingClass, err := client.ProductShippingClass.Get(1, nil)
	if err != nil {
		t.Logf("get shipping class error: %v", err)
	} else {
		t.Logf("got shipping class: id=%d, name=%s", shippingClass.ID, shippingClass.Name)
	}
}

func TestProductShippingClassServiceOp_Update(t *testing.T) {
	shippingClass, err := client.ProductShippingClass.Get(1, nil)
	if shippingClass == nil || err != nil {
		t.Skip("skipping update test: cannot get shipping class 1")
		return
	}
	shippingClass.Name = "Updated Shipping Class"
	res, err := client.ProductShippingClass.Update(shippingClass)
	if err != nil {
		t.Logf("update shipping class error: %v", err)
	} else {
		t.Logf("updated shipping class: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductShippingClassServiceOp_Delete(t *testing.T) {
	shippingClass := ProductShippingClass{
		Name: "Delete Test Shipping Class",
		Slug: "delete-test-shipping-class",
	}
	created, err := client.ProductShippingClass.Create(shippingClass)
	if err != nil {
		t.Skipf("skipping delete test: cannot create shipping class: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.ProductShippingClass.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete shipping class error: %v", err)
	} else {
		t.Logf("deleted shipping class: id=%d", res.ID)
	}
}

func TestProductShippingClassServiceOp_Batch(t *testing.T) {
	data := ProductShippingClassBatchOption{
		Create: []ProductShippingClass{
			{
				Name: "Batch Shipping Class 1",
				Slug: "batch-shipping-class-1",
			},
			{
				Name: "Batch Shipping Class 2",
				Slug: "batch-shipping-class-2",
			},
		},
	}
	res, err := client.ProductShippingClass.Batch(data)
	if err != nil {
		t.Logf("batch shipping classes error: %v", err)
	} else {
		t.Logf("batch created %d shipping classes", len(res.Create))
		for _, sc := range res.Create {
			t.Logf("batch created shipping class: id=%d, name=%s", sc.ID, sc.Name)
		}
	}
}
