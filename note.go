package gowooco

import "time"

// OrderNoteService operate Woo-Commerce Order note, eg: create, view, and delete individual order notes.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-notes
type OrderNoteService interface {
	Create()
	Get()
	List()
	Delete()
}

// OrderNote represent a WooCommerce Order note
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-notes
type OrderNote struct {
	ID             int64      `json:"id,omitempty"`
	Author         string     `json:"author,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DateCreatedGmt *time.Time `json:"date_created_gmt,omitempty"`

	Note         string `json:"note,omitempty"`
	CustomerNote string `json:"customer_note,omitempty"`
	AddedByUser  bool   `json:"added_by_user,omitempty"`
}
