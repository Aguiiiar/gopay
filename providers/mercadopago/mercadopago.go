package mercadopago

import (
	"github.com/Aguiiiar/gopay"
	httpclient "github.com/Aguiiiar/gopay/internal/http_client"
	"github.com/Aguiiiar/gopay/providers/mercadopago/checkoutpro"
)

const Name gopay.ProviderName = "mercadopago"

type Provider struct {
	cfg Config
	hc  *httpclient.Client
}

func New(cfg Config, opts ...Option) (*Provider, error) {

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.mercadopago.com"
	}

	if err := cfg.IsValid(); err != nil {
		return nil, err
	}

	p := &Provider{cfg: cfg}

	for _, opt := range opts {
		opt(p)
	}

	p.hc = httpclient.New(
		cfg.BaseURL,
		httpclient.WithBearerToken(cfg.AccessToken),
	)

	return p, nil
}

type Option func(*Provider)

func (p *Provider) Name() gopay.ProviderName { return Name }

func (p *Provider) CheckoutPro() *checkoutpro.Client {
	return checkoutpro.New(p.hc)
}
