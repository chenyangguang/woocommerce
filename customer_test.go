package woocommerce

import (
	"testing"
	"time"
)

func TestCustomerServiceOp_List(t *testing.T) {
	options := CustomerListOption{
		ListOptions: ListOptions{
			Context: "view",
			Order:   "asc",
			Orderby: "id",
			Page:    1,
			PerPage: 10,
		},
		Role: "customer",
	}
	customers, err := client.Customer.List(options)
	if err != nil {
		t.Logf("error listing customers: %v", err)
	}
	for _, customer := range customers {
		t.Logf("customer: id=%d, email=%s, name=%s %s", customer.ID, customer.Email, customer.FirstName, customer.LastName)
	}
}

func TestCustomerServiceOp_Create(t *testing.T) {
	customer := Customer{
		Email:     "test-customer-" + time.Now().Format("20060102150405") + "@example.com",
		FirstName: "Test",
		LastName:  "Customer",
		Username:  "testuser-" + time.Now().Format("20060102150405"),
		Role:      "customer",
		Billing: CustomerAddress{
			FirstName: "Test",
			LastName:  "Customer",
			Address1:  "123 Test Street",
			City:      "Test City",
			State:     "TS",
			Postcode:  "12345",
			Country:   "US",
			Phone:     "555-1234",
		},
	}
	res, err := client.Customer.Create(customer)
	if err != nil {
		t.Logf("create customer error: %v", err)
	} else {
		t.Logf("created customer: id=%d, email=%s", res.ID, res.Email)
	}
}

func TestCustomerServiceOp_Get(t *testing.T) {
	customer, err := client.Customer.Get(1, nil)
	if err != nil {
		t.Logf("get customer error: %v", err)
	} else {
		t.Logf("got customer: id=%d, email=%s", customer.ID, customer.Email)
	}
}

func TestCustomerServiceOp_Update(t *testing.T) {
	customer, err := client.Customer.Get(1, nil)
	if customer == nil || err != nil {
		t.Skip("skipping update test: cannot get customer 1")
		return
	}
	customer.FirstName = "Updated " + time.Now().Format("20060102150405")
	res, err := client.Customer.Update(customer)
	if err != nil {
		t.Logf("update customer error: %v", err)
	} else {
		t.Logf("updated customer: id=%d, first_name=%s", res.ID, res.FirstName)
	}
}

func TestCustomerServiceOp_Delete(t *testing.T) {
	customer := Customer{
		Email:     "delete-test-" + time.Now().Format("20060102150405") + "@example.com",
		FirstName: "Delete",
		LastName:  "Test",
		Username:  "deletetest-" + time.Now().Format("20060102150405"),
		Role:      "customer",
	}
	created, err := client.Customer.Create(customer)
	if err != nil {
		t.Skipf("skipping delete test: cannot create customer: %v", err)
		return
	}

	optionsDel := DeleteOption{Force: false}
	res, err := client.Customer.Delete(created.ID, optionsDel)
	if err != nil {
		t.Logf("delete customer error: %v", err)
	} else {
		t.Logf("deleted customer: id=%d", res.ID)
	}
}

func TestCustomerServiceOp_Batch(t *testing.T) {
	timeNow := time.Now().Format("20060102150405")
	data := CustomerBatchOption{
		Create: []Customer{
			{
				Email:     "batch1-" + timeNow + "@example.com",
				FirstName: "Batch",
				LastName:  "User 1",
				Username:  "batch1-" + timeNow,
				Role:      "customer",
			},
			{
				Email:     "batch2-" + timeNow + "@example.com",
				FirstName: "Batch",
				LastName:  "User 2",
				Username:  "batch2-" + timeNow,
				Role:      "customer",
			},
		},
	}
	res, err := client.Customer.Batch(data)
	if err != nil {
		t.Logf("batch customers error: %v", err)
	} else {
		t.Logf("batch created %d customers", len(res.Create))
		for _, c := range res.Create {
			t.Logf("batch created customer: id=%d, email=%s", c.ID, c.Email)
		}
	}
}

func TestCustomerServiceOp_GetDownloads(t *testing.T) {
	downloads, err := client.Customer.GetDownloads(1, nil)
	if err != nil {
		t.Logf("get customer downloads error: %v", err)
	} else {
		t.Logf("got %d downloads for customer 1", len(downloads))
		for _, dl := range downloads {
			t.Logf("download: id=%d, name=%s", dl.ID, dl.Name)
		}
	}
}
