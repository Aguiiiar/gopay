package gopay

type ProviderName string

type PaymentProvider interface {
	Name() ProviderName
	ProcessPayment(input ProcessPaymentInput) (ProcessPaymentOutput, error)
}
