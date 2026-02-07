package woocommerce

import (
	"testing"
	"time"
)

func TestProductReviewServiceOp_List(t *testing.T) {
	reviews, err := client.ProductReview.List(nil)
	if err != nil {
		t.Logf("error listing reviews: %v", err)
	}
	for _, review := range reviews {
		t.Logf("review: id=%d, product_id=%d, rating=%d", review.ID, review.ProductID, review.Rating)
	}
}

func TestProductReviewServiceOp_Create(t *testing.T) {
	review := ProductReview{
		ProductID: 1,
		Review:    "Test review " + time.Now().Format("20060102150405"),
		Reviewer:  "Test Reviewer",
		Rating:    5,
		Verified:  true,
	}
	res, err := client.ProductReview.Create(review)
	if err != nil {
		t.Logf("create review error: %v", err)
	} else {
		t.Logf("created review: id=%d, rating=%d", res.ID, res.Rating)
	}
}

func TestProductReviewServiceOp_Get(t *testing.T) {
	review, err := client.ProductReview.Get(1, nil)
	if err != nil {
		t.Logf("get review error: %v", err)
	} else {
		t.Logf("got review: id=%d, rating=%d", review.ID, review.Rating)
	}
}

func TestProductReviewServiceOp_Update(t *testing.T) {
	review, err := client.ProductReview.Get(1, nil)
	if review == nil || err != nil {
		t.Skip("skipping update test: cannot get review 1")
		return
	}
	review.Review = "Updated review " + time.Now().Format("20060102150405")
	res, err := client.ProductReview.Update(review)
	if err != nil {
		t.Logf("update review error: %v", err)
	} else {
		t.Logf("updated review: id=%d, review=%s", res.ID, res.Review)
	}
}

func TestProductReviewServiceOp_Delete(t *testing.T) {
	review := ProductReview{
		ProductID: 1,
		Review:    "Test review to delete",
		Reviewer:  "Test Reviewer",
		Rating:    3,
		Verified:  true,
	}
	created, err := client.ProductReview.Create(review)
	if err != nil {
		t.Skipf("skipping delete test: cannot create review: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.ProductReview.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete review error: %v", err)
	} else {
		t.Logf("deleted review: id=%d", res.ID)
	}
}

func TestProductReviewServiceOp_Batch(t *testing.T) {
	timeNow := time.Now().Format("20060102150405")
	data := ProductReviewBatchOption{
		Create: []ProductReview{
			{
				ProductID: 1,
				Review:    "Batch review 1 " + timeNow,
				Reviewer:  "Batch Reviewer 1",
				Rating:    5,
			},
			{
				ProductID: 1,
				Review:    "Batch review 2 " + timeNow,
				Reviewer:  "Batch Reviewer 2",
				Rating:    4,
			},
		},
	}
	res, err := client.ProductReview.Batch(data)
	if err != nil {
		t.Logf("batch reviews error: %v", err)
	} else {
		t.Logf("batch created %d reviews", len(res.Create))
		for _, r := range res.Create {
			t.Logf("batch created review: id=%d, rating=%d", r.ID, r.Rating)
		}
	}
}
