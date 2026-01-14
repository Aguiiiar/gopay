//go:build integration

package mercadopago_test

import (
	"context"
	"os"
	"testing"

	"github.com/Aguiiiar/gopay/providers/mercadopago"
)

func TestIntegration_CreatePreference_Real(t *testing.T) {
	token := os.Getenv("MP_ACCESS_TOKEN")
	if token == "" {
		t.Skip("MP_ACCESS_TOKEN not set")
	}

	baseURL := os.Getenv("MP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.mercadopago.com"
	}

	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     baseURL,
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	out, err := mp.CreatePreference(context.Background(), mercadopago.CreatePreferenceInput{
		Items: []mercadopago.PreferenceItem{
			{
				Title:      "GoPay Integration Test",
				Quantity:   1,
				CurrencyID: "BRL",
				UnitPrice:  1.00,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreatePreference error: %v", err)
	}

	// 👇 logs de sucesso
	t.Log("Mercado Pago preference created successfully")
	t.Logf("Preference ID: %s", out.ID)
	t.Logf("InitPoint: %s", out.InitPoint)
	t.Logf("SandboxInitPoint: %s", out.SandboxInitPoint)

	if out.ID == "" {
		t.Fatalf("expected preference id, got empty")
	}
	if out.InitPoint == "" && out.SandboxInitPoint == "" {
		t.Fatalf("expected init_point or sandbox_init_point, got empty")
	}
}
