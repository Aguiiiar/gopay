package gopay

import "fmt"

type Client struct {
	providers map[ProviderName]PaymentProvider
}

func New(cfg Config, opts ...Option) (*Client, error) {
	c := &Client{
		providers: map[ProviderName]PaymentProvider{},
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	if len(c.providers) == 0 {
		return nil, ErrNoProvidersEnabled
	}

	return c, nil
}

func (c *Client) Provider(name ProviderName) (PaymentProvider, error) {
	p, ok := c.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider not configured: %s", name)
	}
	return p, nil
}
