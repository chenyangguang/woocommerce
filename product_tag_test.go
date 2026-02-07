package woocommerce

import (
	"testing"
)

func TestProductTagServiceOp_List(t *testing.T) {
	tags, err := client.ProductTag.List(nil)
	if err != nil {
		t.Logf("error listing tags: %v", err)
	}
	for _, tag := range tags {
		t.Logf("tag: id=%d, name=%s, slug=%s", tag.ID, tag.Name, tag.Slug)
	}
}

func TestProductTagServiceOp_Create(t *testing.T) {
	tag := ProductTag{
		Name: "Test Tag",
		Slug: "test-tag",
	}
	res, err := client.ProductTag.Create(tag)
	if err != nil {
		t.Logf("create tag error: %v", err)
	} else {
		t.Logf("created tag: id=%d, name=%s", res.ID, res.Name)
	}
}

func TestProductTagServiceOp_Get(t *testing.T) {
	tag, err := client.ProductTag.Get(1, nil)
	if err != nil {
		t.Logf("get tag error: %v", err)
	} else {
		t.Logf("got tag: id=%d, name=%s", tag.ID, tag.Name)
	}
}

func TestProductTagServiceOp_Update(t *testing.T) {
	tag, err := client.ProductTag.Get(1, nil)
	if tag == nil || err != nil {
		t.Skip("skipping update test: cannot get tag 1")
		return
	}
	tag.Description = "Updated description"
	res, err := client.ProductTag.Update(tag)
	if err != nil {
		t.Logf("update tag error: %v", err)
	} else {
		t.Logf("updated tag: id=%d, description=%s", res.ID, res.Description)
	}
}

func TestProductTagServiceOp_Delete(t *testing.T) {
	tag := ProductTag{
		Name: "Delete Test Tag",
		Slug: "delete-test-tag",
	}
	created, err := client.ProductTag.Create(tag)
	if err != nil {
		t.Skipf("skipping delete test: cannot create tag: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.ProductTag.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete tag error: %v", err)
	} else {
		t.Logf("deleted tag: id=%d", res.ID)
	}
}

func TestProductTagServiceOp_Batch(t *testing.T) {
	data := ProductTagBatchOption{
		Create: []ProductTag{
			{
				Name: "Batch Tag 1",
				Slug: "batch-tag-1",
			},
			{
				Name: "Batch Tag 2",
				Slug: "batch-tag-2",
			},
		},
	}
	res, err := client.ProductTag.Batch(data)
	if err != nil {
		t.Logf("batch tags error: %v", err)
	} else {
		t.Logf("batch created %d tags", len(res.Create))
		for _, tag := range res.Create {
			t.Logf("batch created tag: id=%d, name=%s", tag.ID, tag.Name)
		}
	}
}
