package mercadopago

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aguiiiar/gopay"
	httpclient "github.com/Aguiiiar/gopay/internal/http_client"
)

const Name gopay.ProviderName = "mercadopago"

type Provider struct {
	cfg      Config
	httpBase *http.Client
	hc       *httpclient.Client
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

	hc := httpclient.New(
		cfg.BaseURL,
		httpclient.WithBearerToken(cfg.AccessToken),
	)

	return &Provider{hc: hc}, nil
}

type Option func(*Provider)

func (p *Provider) Name() gopay.ProviderName { return Name }

// func WithHTTPClient(httpClient *http.Client) Option {
// 	return func(p *Provider) {
// 		if httpClient != nil {
// 			p.httpBase = httpClient
// 		}
// 	}
// }

func (p *Provider) CreatePreference(ctx context.Context, input CreatePreferenceInput) (CreatePreferenceOutput, error) {
	response, err := p.hc.DoJSON(ctx, http.MethodPost, "/checkout/preferences", input, nil)
	if err != nil {
		return CreatePreferenceOutput{}, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return CreatePreferenceOutput{}, fmt.Errorf("MercadoPago: api status=%d body=%s", response.StatusCode, string(response.Body))
	}

	var output CreatePreferenceOutput
	if err := json.Unmarshal(response.Body, &output); err != nil {
		return CreatePreferenceOutput{}, fmt.Errorf("MercadoPago: unmarshal create preference response: %w", err)
	}

	return output, nil
}
