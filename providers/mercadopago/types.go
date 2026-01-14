package mercadopago

type CreatePreferenceInput struct {
	Items []PreferenceItem `json:"items"`
}

type PreferenceItem struct {
	Title      string  `json:"title"`
	Quantity   int     `json:"quantity"`
	CurrencyID string  `json:"currency_id"`
	UnitPrice  float64 `json:"unit_price"`
}

type CreatePreferenceOutput struct {
	ID               string `json:"id"`
	InitPoint        string `json:"init_point"`
	SandboxInitPoint string `json:"sandbox_init_point"`
}
