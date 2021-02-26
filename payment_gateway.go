package woocommerce

import "fmt"

const (
	paymentGatewayBasePath = "payment_gateways"
)

// PaymentGatewayService is an interface for interfacing with the payment-gateways endpoints of woocommerce API
// https://woocommerce.github.io/woocommerce-rest-api-docs/#payment-gateways
type PaymentGatewayService interface {
	Get(id string) (*PaymentGateway, error)
	List(options interface{}) ([]PaymentGateway, error)
	Update(pg *PaymentGateway) (*PaymentGateway, error)
}

// PaymentGatewayServiceOp handles communication with the payment gateway related methods of WooCommerce restful api
type PaymentGatewayServiceOp struct {
	client *Client
}

// PaymentGateway represents a woocommerce payment_gateway
// https://woocommerce.github.io/woocommerce-rest-api-docs/#payment-gateway-properties
type PaymentGateway struct {
	ID                string   `json:"id,omitempty"`
	Title             string   `json:"title,omitempty"`
	Description       string   `json:"description,omitempty"`
	Order             string   `json:"order,omitempty"`
	Enabled           bool     `json:"enabled,omitempty"`
	MethodTitle       string   `json:"method_title,omitempty"`
	MethodDescription string   `json:"method_description,omitempty"`
	MethodSupports    []string `json:"method_supports,omitempty"`
	Settings          *Setting `json:"settings,omitempty"`
	Links             Links    `json:"_links,omitempty"`
}

type Setting struct {
	Title            *PaymentSetting `json:"title,omitempty"`
	Email            *PaymentSetting `json:"email,omitempty"`
	Advanced         *PaymentSetting `json:"advanced,omitempty"`
	Instructions     *PaymentSetting `json:"instructions,omitempty"`
	TestMode         *PaymentSetting `json:"testmode,omitempty"`
	Debug            *PaymentSetting `json:"debug,omitempty"`
	ImageUrl         *PaymentSetting `json:"image_url,omitempty"`
	PageStyle        *PaymentSetting `json:"page_style,omitempty"`
	Paymentaction    *PaymentSetting `json:"paymentaction,omitempty"`
	AddressOverride  *PaymentSetting `json:"address_override,omitempty"`
	SendShipping     *PaymentSetting `json:"send_shipping,omitempty"`
	InvoicePrefix    *PaymentSetting `json:"invoice_prefix,omitempty"`
	IdentityToken    *PaymentSetting `json:"identity_token,omitempty"`
	ReceiverEmail    *PaymentSetting `json:"receiver_email,omitempty"`
	IpnNotification  *PaymentSetting `json:"ipn_notification,omitempty"`
	EnableForMethods *PaymentSetting `json:"enable_for_methods,omitempty"`
	EnableForVirtual *PaymentSetting `json:"enable_for_virtual,omitempty"`
}

// PaymentSetting  represents a WooCommerce payment-gateway
// https://woocommerce.github.io/woocommerce-rest-api-docs/#payment-gateway-settings-properties
// id	string	A unique identifier for the setting.READ-ONLY
// label	string	A human readable label for the setting used in interfaces.READ-ONLY
// description	string	A human readable description for the setting used in interfaces.READ-ONLY
// type	string	Type of setting. Options: text, email, number, color, password, textarea, select, multiselect, radio, image_width and checkbox.READ-ONLY
// value	string	Setting value.
// default	string	Default value for the setting.READ-ONLY
// tip	string	Additional help text shown to the user about the setting.READ-ONLY
// placeholder	string	Placeholder text to be displayed in text inputs.
type PaymentSetting struct {
	ID          string `json:"id,omitempty"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	Default     string `json:"default,omitempty"`
	Tip         string `json:"tip,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
}

// List return multiple payment gateway
// https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-payment-gateways
func (p *PaymentGatewayServiceOp) List(options interface{}) ([]PaymentGateway, error) {
	path := fmt.Sprintf("%s", paymentGatewayBasePath)
	resource := make([]PaymentGateway, 0)
	err := p.client.Get(path, &resource, options)
	return resource, err
}

// Get implement for retrieve and view a specific payment gateway
// https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-an-payment-gateway
func (p *PaymentGatewayServiceOp) Get(id string) (*PaymentGateway, error) {
	path := fmt.Sprintf("%s/%s", paymentGatewayBasePath, id)
	resource := new(PaymentGateway)
	err := p.client.Get(path, &resource, nil)
	return resource, err
}

// Update method allow you to make changes to a payment gateway
// https://woocommerce.github.io/woocommerce-rest-api-docs/#update-a-payment-gateway
func (p *PaymentGatewayServiceOp) Update(pg *PaymentGateway) (*PaymentGateway, error) {
	path := fmt.Sprintf("%s/%s", paymentGatewayBasePath, pg.ID)
	resource := new(PaymentGateway)
	err := p.client.Put(path, pg, &resource)

	return resource, err
}
