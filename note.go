package woocommerce

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

func (n *OrderNoteServiceOp) Create() {}
func (n *OrderNoteServiceOp) Get()    {}
func (n *OrderNoteServiceOp) List()   {}
func (n *OrderNoteServiceOp) Delete() {}
