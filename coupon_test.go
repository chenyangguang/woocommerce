package woocommerce

import (
	"testing"
	"time"
)

func TestCouponServiceOp_List(t *testing.T) {
	options := CouponListOption{
		ListOptions: ListOptions{
			Context: "view",
			Order:   "desc",
			Orderby: "date",
			Page:    1,
			PerPage: 10,
		},
	}
	coupons, err := client.Coupon.List(options)
	if err != nil {
		t.Logf("error listing coupons: %v", err)
	}
	for _, coupon := range coupons {
		t.Logf("coupon: id=%d, code=%s, amount=%s", coupon.ID, coupon.Code, coupon.Amount)
	}
}

func TestCouponServiceOp_Create(t *testing.T) {
	coupon := Coupon{
		Code:          "TEST" + time.Now().Format("20060102150405"),
		DiscountType:  "fixed_cart",
		Amount:        "10.00",
		Description:   "Test coupon",
		UsageCount:    0,
		FreeShipping:  false,
		IndividualUse: false,
	}
	res, err := client.Coupon.Create(coupon)
	if err != nil {
		t.Logf("create coupon error: %v", err)
	} else {
		t.Logf("created coupon: id=%d, code=%s", res.ID, res.Code)
	}
}

func TestCouponServiceOp_Get(t *testing.T) {
	coupon, err := client.Coupon.Get(1, nil)
	if err != nil {
		t.Logf("get coupon error: %v", err)
	} else {
		t.Logf("got coupon: id=%d, code=%s", coupon.ID, coupon.Code)
	}
}

func TestCouponServiceOp_Update(t *testing.T) {
	coupon, err := client.Coupon.Get(1, nil)
	if coupon == nil || err != nil {
		t.Skip("skipping update test: cannot get coupon 1")
		return
	}
	coupon.Description = "Updated description " + time.Now().Format("20060102150405")
	res, err := client.Coupon.Update(coupon)
	if err != nil {
		t.Logf("update coupon error: %v", err)
	} else {
		t.Logf("updated coupon: id=%d, description=%s", res.ID, res.Description)
	}
}

func TestCouponServiceOp_Delete(t *testing.T) {
	coupon := Coupon{
		Code:         "DELETE" + time.Now().Format("20060102150405"),
		DiscountType: "fixed_cart",
		Amount:       "5.00",
		Description:  "Test coupon to delete",
	}
	created, err := client.Coupon.Create(coupon)
	if err != nil {
		t.Skipf("skipping delete test: cannot create coupon: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.Coupon.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete coupon error: %v", err)
	} else {
		t.Logf("deleted coupon: id=%d", res.ID)
	}
}

func TestCouponServiceOp_Batch(t *testing.T) {
	timeNow := time.Now().Format("20060102150405")
	data := CouponBatchOption{
		Create: []Coupon{
			{
				Code:         "BATCH1" + timeNow,
				DiscountType: "fixed_cart",
				Amount:       "10.00",
				Description:  "Batch coupon 1",
			},
			{
				Code:         "BATCH2" + timeNow,
				DiscountType: "percent",
				Amount:       "15",
				Description:  "Batch coupon 2",
			},
		},
	}
	res, err := client.Coupon.Batch(data)
	if err != nil {
		t.Logf("batch coupons error: %v", err)
	} else {
		t.Logf("batch created %d coupons", len(res.Create))
		for _, c := range res.Create {
			t.Logf("batch created coupon: id=%d, code=%s", c.ID, c.Code)
		}
	}
}
