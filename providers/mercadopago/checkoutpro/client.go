package checkoutpro

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httpclient "github.com/Aguiiiar/gopay/internal/http_client"
	"github.com/Aguiiiar/gopay/providers/mercadopago/mperrors"
)

type Client struct {
	hc *httpclient.Client
}

func New(hc *httpclient.Client) *Client {
	return &Client{hc: hc}
}

// CreatePreference creates a Checkout Pro preference in Mercado Pago.
//
// API reference:
// https://www.mercadopago.com.br/developers/en/reference/preferences/_checkout_preferences/post
//
// This method uses Checkout Pro flow.
// It returns URLs for redirecting the user to Mercado Pago.
func (p *Client) CreatePreference(ctx context.Context, input CreatePreferenceInput) (CreatePreferenceOutput, error) {
	response, err := p.hc.DoJSON(ctx, http.MethodPost, "/checkout/preferences", input, nil)
	if err != nil {
		return CreatePreferenceOutput{}, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return CreatePreferenceOutput{}, mperrors.Parse(response.StatusCode, response.Body)
	}

	var output CreatePreferenceOutput
	if err := json.Unmarshal(response.Body, &output); err != nil {
		return CreatePreferenceOutput{}, fmt.Errorf("MercadoPago: unmarshal create preference response: %w", err)
	}

	return output, nil
}

// GetPreference retrieves a Checkout Pro preference from Mercado Pago by its ID.
//
// API reference:
// https://www.mercadopago.com.br/developers/en/reference/preferences/_checkout_preferences/{id}/get
func (p *Client) GetPreference(ctx context.Context, preferenceID string) (GetPreferenceOutput, error) {
	path := fmt.Sprintf("/checkout/preferences/%s", preferenceID)
	response, err := p.hc.DoJSON(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return GetPreferenceOutput{}, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return GetPreferenceOutput{}, mperrors.Parse(response.StatusCode, response.Body)
	}

	var output GetPreferenceOutput
	if err := json.Unmarshal(response.Body, &output); err != nil {
		return GetPreferenceOutput{}, fmt.Errorf("MercadoPago: unmarshal get preference response: %w", err)
	}

	return output, nil
}

// UpdatePreference updates an existing Checkout Pro preference in Mercado Pago.
//
// API reference:
// https://www.mercadopago.com.br/developers/en/reference/preferences/_checkout_preferences/{id}/put
func (p *Client) UpdatePreference(ctx context.Context, preferenceID string, input UpdatePreferenceInput) (UpdatePreferenceOutput, error) {
	path := fmt.Sprintf("/checkout/preferences/%s", preferenceID)
	response, err := p.hc.DoJSON(ctx, http.MethodPut, path, input, nil)
	if err != nil {
		return UpdatePreferenceOutput{}, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return UpdatePreferenceOutput{}, mperrors.Parse(response.StatusCode, response.Body)
	}

	var output UpdatePreferenceOutput
	if err := json.Unmarshal(response.Body, &output); err != nil {
		return UpdatePreferenceOutput{}, fmt.Errorf("MercadoPago: unmarshal update preference response: %w", err)
	}

	return output, nil
}
