package woocommerce

import "fmt"

const (
	orderNoteBasePath = "orders"
)

// OrderNoteService operate Woo-Commerce Order note, eg: create, view, and delete individual order notes.
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-notes
type OrderNoteService interface {
	Create(text string) (*OrderNote, error)
	Get(orderId int64, noteId int64) (*OrderNote, error)
	List(orderId int64, options interface{})
	Delete()
}

// OrderNote represent a WooCommerce Order note
// https://woocommerce.github.io/woocommerce-rest-api-docs/#order-notes
type OrderNote struct {
	ID             int64  `json:"id,omitempty"`
	Author         string `json:"author,omitempty"`
	DateCreated    string `json:"date_created,omitempty"`
	DateCreatedGmt string `json:"date_created_gmt,omitempty"`

	Note         string `json:"note,omitempty"`
	CustomerNote string `json:"customer_note,omitempty"`
	AddedByUser  bool   `json:"added_by_user,omitempty"`
}

type OrderNoteServiceOp struct {
	client *Client
}

func (n *OrderNoteServiceOp) Create(orderId int64, text string) (*OrderNote, error) {
	path := fmt.Sprintf("%s/%d/notes", orderNoteBasePath, orderId)
	resource := new(OrderNote)
	insertOrderNote := OrderNote{
		Note: text,
	}
	err := n.client.Post(path, insertOrderNote, resource)
	return resource, err
}

func (n *OrderNoteServiceOp) Get(orderId int64, noteId int64) (*OrderNote, error) {
	path := fmt.Sprintf("%s/%d/notes/%d", orderNoteBasePath, orderId, noteId)
	resource := new(OrderNote)

	err := n.client.Get(path, resource, nil)
	return resource, err
}

func (n *OrderNoteServiceOp) List(orderId int64, optoions interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s/%d/notes", orderNoteBasePath, orderId)
	resource := new(OrderNote)

	err := n.client.Get(path, resource, options)
	return resource, err
}

func (n *OrderNoteServiceOp) Delete(orderId int64, noteId int64, options interface{}) (*OrderNote, error) {
	path := fmt.Sprintf("%s/%d/notes/%d", orderNoteBasePath, orderId, noteId)
	resource := new(OrderNote)
	err := n.client.Delete(path, options, &resource)
	return resource, err
}
