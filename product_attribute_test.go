package woocommerce

import (
	"testing"
)

func TestProductAttributeServiceOp_List(t *testing.T) {
	attributes, err := client.ProductAttribute.List(nil)
	if err != nil {
		t.Logf("error listing attributes: %v", err)
	}
	for _, attr := range attributes {
		t.Logf("attribute: id=%d, name=%s, slug=%s", attr.ID, attr.Name, attr.Slug)
	}
}

func TestProductAttributeServiceOp_Create(t *testing.T) {
	attribute := ProductAttributeData{
		Name: "Test Attribute",
		Slug: "test-attribute",
		Type: "select",
	}
	res, err := client.ProductAttribute.Create(attribute)
	if err != nil {
		t.Logf("create attribute error: %v", err)
	} else {
		t.Logf("created attribute: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductAttributeServiceOp_Get(t *testing.T) {
	attribute, err := client.ProductAttribute.Get(1, nil)
	if err != nil {
		t.Logf("get attribute error: %v", err)
	} else {
		t.Logf("got attribute: id=%d, name=%s", attribute.ID, attribute.Name)
	}
}

func TestProductAttributeServiceOp_Update(t *testing.T) {
	attribute, err := client.ProductAttribute.Get(1, nil)
	if attribute == nil || err != nil {
		t.Skip("skipping update test: cannot get attribute 1")
		return
	}
	attribute.Visible = false
	res, err := client.ProductAttribute.Update(attribute)
	if err != nil {
		t.Logf("update attribute error: %v", err)
	} else {
		t.Logf("updated attribute: id=%d, visible=%v", res.ID, res.Visible)
	}
}

func TestProductAttributeServiceOp_Delete(t *testing.T) {
	attribute := ProductAttributeData{
		Name: "Delete Test Attribute",
		Slug: "delete-test-attribute",
		Type: "text",
	}
	created, err := client.ProductAttribute.Create(attribute)
	if err != nil {
		t.Skipf("skipping delete test: cannot create attribute: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.ProductAttribute.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete attribute error: %v", err)
	} else {
		t.Logf("deleted attribute: id=%d", res.ID)
	}
}

func TestProductAttributeServiceOp_Batch(t *testing.T) {
	data := ProductAttributeBatchOption{
		Create: []ProductAttributeData{
			{
				Name: "Batch Attribute 1",
				Slug: "batch-attribute-1",
				Type: "select",
			},
			{
				Name: "Batch Attribute 2",
				Slug: "batch-attribute-2",
				Type: "text",
			},
		},
	}
	res, err := client.ProductAttribute.Batch(data)
	if err != nil {
		t.Logf("batch attributes error: %v", err)
	} else {
		t.Logf("batch created %d attributes", len(res.Create))
		for _, a := range res.Create {
			t.Logf("batch created attribute: id=%d, name=%s", a.ID, a.Name)
		}
	}
}
