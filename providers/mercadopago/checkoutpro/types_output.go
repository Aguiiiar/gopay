package checkoutpro

type CreatePreferenceOutput struct {
	ID               string `json:"id,omitempty"`
	InitPoint        string `json:"init_point"`
	SandboxInitPoint string `json:"sandbox_init_point"`

	AutoReturn        string    `json:"auto_return,omitempty"`
	BackURLs          *BackURLs `json:"back_urls,omitempty"`
	NotificationURL   string    `json:"notification_url,omitempty"`
	ExternalReference string    `json:"external_reference,omitempty"`

	DateCreated        string `json:"date_created,omitempty"`
	ExpirationDateFrom string `json:"expiration_date_from,omitempty"`
	ExpirationDateTo   string `json:"expiration_date_to,omitempty"`
	Expires            bool   `json:"expires,omitempty"`

	ClientID    string `json:"client_id,omitempty"`
	CollectorID int64  `json:"collector_id,omitempty"`
	SiteID      string `json:"site_id,omitempty"`

	Items []PreferenceItemOutput `json:"items,omitempty"`
	Payer *PayerOutput           `json:"payer,omitempty"`

	PaymentMethods *PaymentMethodsOutput `json:"payment_methods,omitempty"`

	Metadata         map[string]any    `json:"metadata,omitempty"`
	AdditionalInfo   any               `json:"additional_info,omitempty"`
	InternalMetadata any               `json:"internal_metadata,omitempty"`
	ProcessingModes  any               `json:"processing_modes,omitempty"`
	ProductID        any               `json:"product_id,omitempty"`
	RedirectURLs     map[string]string `json:"redirect_urls,omitempty"`
	Shipments        any               `json:"shipments,omitempty"`
	TotalAmount      any               `json:"total_amount,omitempty"`
	LastUpdated      any               `json:"last_updated,omitempty"`
	OperationType    string            `json:"operation_type,omitempty"`
	Marketplace      string            `json:"marketplace,omitempty"`
	MarketplaceFee   float64           `json:"marketplace_fee,omitempty"`
}

type GetPreferenceOutput struct {
	ID               string `json:"id,omitempty"`
	InitPoint        string `json:"init_point"`
	SandboxInitPoint string `json:"sandbox_init_point"`

	AutoReturn        string    `json:"auto_return,omitempty"`
	BackURLs          *BackURLs `json:"back_urls,omitempty"`
	NotificationURL   string    `json:"notification_url,omitempty"`
	ExternalReference string    `json:"external_reference,omitempty"`

	DateCreated        string `json:"date_created,omitempty"`
	ExpirationDateFrom string `json:"expiration_date_from,omitempty"`
	ExpirationDateTo   string `json:"expiration_date_to,omitempty"`
	Expires            bool   `json:"expires,omitempty"`

	ClientID    string `json:"client_id,omitempty"`
	CollectorID int64  `json:"collector_id,omitempty"`
	SiteID      string `json:"site_id,omitempty"`

	Items []PreferenceItemOutput `json:"items,omitempty"`
	Payer *PayerOutput           `json:"payer,omitempty"`

	PaymentMethods *PaymentMethodsOutput `json:"payment_methods,omitempty"`
}

type PreferenceItemOutput struct {
	ID          string     `json:"id,omitempty"`
	CategoryID  string     `json:"category_id,omitempty"`
	CurrencyID  CurrencyID `json:"currency_id,omitempty"`
	Description string     `json:"description,omitempty"`
	Title       string     `json:"title,omitempty"`
	Quantity    int        `json:"quantity,omitempty"`
	UnitPrice   float64    `json:"unit_price,omitempty"`
}

type PayerOutput struct {
	Phone          *Phone          `json:"phone,omitempty"`
	Address        *Address        `json:"address,omitempty"`
	Email          string          `json:"email,omitempty"`
	Identification *Identification `json:"identification,omitempty"`
	Name           string          `json:"name,omitempty"`
	Surname        string          `json:"surname,omitempty"`
	DateCreated    any             `json:"date_created,omitempty"`
	LastPurchase   any             `json:"last_purchase,omitempty"`
}

type PaymentMethodsOutput struct {
	DefaultCardID          any        `json:"default_card_id,omitempty"`
	DefaultPaymentMethodID string     `json:"default_payment_method_id,omitempty"`
	ExcludedPaymentMethods []IDObject `json:"excluded_payment_methods,omitempty"`
	ExcludedPaymentTypes   []IDObject `json:"excluded_payment_types,omitempty"`
	Installments           int        `json:"installments,omitempty"`
	DefaultInstallments    any        `json:"default_installments,omitempty"`
}
