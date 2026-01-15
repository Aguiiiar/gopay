package checkoutpro

type IDObject struct {
	ID string `json:"id,omitempty"`
}

type CurrencyID string

const (
	CurrencyIDARS CurrencyID = "ARS"
	CurrencyIDBRL CurrencyID = "BRL"
	CurrencyIDCLP CurrencyID = "CLP"
	CurrencyIDMXN CurrencyID = "MXN"
	CurrencyIDCOP CurrencyID = "COP"
	CurrencyIDPEN CurrencyID = "PEN"
	CurrencyIDUYU CurrencyID = "UYU"
)

type AutoReturn string

const (
	AutoReturnApproved AutoReturn = "approved"
	AutoReturnAll      AutoReturn = "all"
)
