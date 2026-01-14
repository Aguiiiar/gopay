package gopay

type Option func(*Client) error

func WithProvider(p PaymentProvider) Option {
	return func(c *Client) error {
		c.providers[p.Name()] = p
		return nil
	}
}
