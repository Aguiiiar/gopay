package checkoutpro

type CreatePreferenceInput struct {
	Items               []PreferenceItem `json:"items"`
	BackURLs            *BackURLs        `json:"back_urls,omitempty"`
	AutoReturn          string           `json:"auto_return,omitempty"`
	StatementDescriptor string           `json:"statement_descriptor,omitempty"`
	ExternalReference   string           `json:"external_reference,omitempty"`
	NotificationURL     string           `json:"notification_url,omitempty"`
	Expires             bool             `json:"expires,omitempty"`
	ExpirationDateFrom  string           `json:"expiration_date_from,omitempty"`
	ExpirationDateTo    string           `json:"expiration_date_to,omitempty"`
	PaymentMethods      *PaymentMethods  `json:"payment_methods"`
	Payer               *Payers          `json:"payer"`
	Shipments           *Shipments       `json:"shipments,omitempty"`
}

type PreferenceItem struct {
	ID          string      `json:"id,omitempty"`
	Title       string      `json:"title"`
	Quantity    int         `json:"quantity"`
	UnitPrice   float64     `json:"unit_price"`
	CurrencyID  *CurrencyID `json:"currency_id"`
	Description *string     `json:"description,omitempty"`
	PictureURL  *string     `json:"picture_url,omitempty"`
}

type BackURLs struct {
	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`
	Pending string `json:"pending,omitempty"`
}

type Payers struct {
	Email          string         `json:"email"`
	Name           string         `json:"name,omitempty"`
	Surname        string         `json:"surname,omitempty"`
	Phone          Phone          `json:"phone"`
	Identification Identification `json:"identification"`
	Address        Address        `json:"address"`
}

type PaymentMethods struct {
	ExcludedPaymentMethods []any  `json:"excluded_payment_methods,omitempty"`
	ExcludedPaymentTypes   []any  `json:"excluded_payment_types,omitempty"`
	Installments           int    `json:"installments,omitempty"`
	DefaultPaymentMethodID string `json:"default_payment_method_id,omitempty"`
}

type Phone struct {
	AreaCode string `json:"area_code,omitempty"`
	Number   string `json:"number,omitempty"`
}

type Identification struct {
	Type   string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type Address struct {
	StreetName   string `json:"street_name,omitempty"`
	StreetNumber int    `json:"street_number,omitempty"`
	ZipCode      string `json:"zip_code,omitempty"`
}

type Shipments struct {
	LocalPickup           bool            `json:"local_pickup,omitempty"`
	Dimensions            string          `json:"dimensions,omitempty"`
	DefaultShippingMethod any             `json:"default_shipping_method,omitempty"`
	FreeMethods           []any           `json:"free_methods,omitempty"`
	Cost                  float64         `json:"cost,omitempty"`
	FreeShipping          bool            `json:"free_shipping,omitempty"`
	ReceiverAddress       ReceiverAddress `json:"receiver_address,omitempty"`
}

type ReceiverAddress struct {
	ZipCode      string `json:"zip_code,omitempty"`
	StreetName   string `json:"street_name,omitempty"`
	CityName     string `json:"city_name,omitempty"`
	StateName    string `json:"state_name,omitempty"`
	StreetNumber int    `json:"street_number,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
}
