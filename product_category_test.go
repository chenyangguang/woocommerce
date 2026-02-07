package woocommerce

import (
	"testing"
)

func TestProductCategoryServiceOp_List(t *testing.T) {
	categories, err := client.ProductCategory.List(nil)
	if err != nil {
		t.Logf("error listing categories: %v", err)
	}
	for _, category := range categories {
		t.Logf("category: id=%d, name=%s, slug=%s", category.ID, category.Name, category.Slug)
	}
}

func TestProductCategoryServiceOp_Create(t *testing.T) {
	category := ProductCategory{
		Name: "Test Category",
		Slug: "test-category",
	}
	res, err := client.ProductCategory.Create(category)
	if err != nil {
		t.Logf("create category error: %v", err)
	} else {
		t.Logf("created category: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductCategoryServiceOp_Get(t *testing.T) {
	category, err := client.ProductCategory.Get(1, nil)
	if err != nil {
		t.Logf("get category error: %v", err)
	} else {
		t.Logf("got category: id=%d, name=%s", category.ID, category.Name)
	}
}

func TestProductCategoryServiceOp_Update(t *testing.T) {
	category, err := client.ProductCategory.Get(1, nil)
	if category == nil || err != nil {
		t.Skip("skipping update test: cannot get category 1")
		return
	}
	category.Description = "Updated description"
	res, err := client.ProductCategory.Update(category)
	if err != nil {
		t.Logf("update category error: %v", err)
	} else {
		t.Logf("updated category: id=%d, description=%s", res.ID, res.Description)
	}
}

func TestProductCategoryServiceOp_Delete(t *testing.T) {
	category := ProductCategory{
		Name: "Delete Test Category",
		Slug: "delete-test-category",
	}
	created, err := client.ProductCategory.Create(category)
	if err != nil {
		t.Skipf("skipping delete test: cannot create category: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.ProductCategory.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete category error: %v", err)
	} else {
		t.Logf("deleted category: id=%d", res.ID)
	}
}

func TestProductCategoryServiceOp_Batch(t *testing.T) {
	data := ProductCategoryBatchOption{
		Create: []ProductCategory{
			{
				Name: "Batch Category 1",
				Slug: "batch-category-1",
			},
			{
				Name: "Batch Category 2",
				Slug: "batch-category-2",
			},
		},
	}
	res, err := client.ProductCategory.Batch(data)
	if err != nil {
		t.Logf("batch categories error: %v", err)
	} else {
		t.Logf("batch created %d categories", len(res.Create))
		for _, c := range res.Create {
			t.Logf("batch created category: id=%d, name=%s", c.ID, c.Name)
		}
	}
}
