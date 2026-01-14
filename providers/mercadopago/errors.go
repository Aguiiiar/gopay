package mercadopago

import "errors"

var (
	ErrInvalidConfig  = errors.New("invalid mercadopago config")
	ErrNotImplemented = errors.New("mercadopago: not implemented yet")
)
