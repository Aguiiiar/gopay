package gopay

type ProcessPaymentInput struct {
	Amount   int64
	Currency string
	OrderID  string
}

type ProcessPaymentOutput struct {
	ProviderPaymentID string
	Status            string
}
