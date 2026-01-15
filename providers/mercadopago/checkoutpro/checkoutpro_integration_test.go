//go:build integration
// +build integration

package checkoutpro_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Aguiiiar/gopay/providers/mercadopago"
	"github.com/Aguiiiar/gopay/providers/mercadopago/checkoutpro"
)

func TestCheckoutProIntegration_CreatePreference_Real(t *testing.T) {
	token := os.Getenv("MP_ACCESS_TOKEN")
	if token == "" {
		t.Skip("MP_ACCESS_TOKEN not set")
	}

	baseURL := os.Getenv("MP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.mercadopago.com"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     baseURL,
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	cid := checkoutpro.CurrencyIDBRL

	out, err := mp.CheckoutPro().CreatePreference(ctx, checkoutpro.CreatePreferenceInput{
		Items: []checkoutpro.PreferenceItem{
			{
				Title:      "GoPay Integration Test - Checkout Pro",
				Quantity:   1,
				UnitPrice:  1.00,
				CurrencyID: &cid,
			},
		},
		ExternalReference: "gopay-integration-test",
		AutoReturn:        "approved",
		BackURLs: &checkoutpro.BackURLs{
			Success: "https://example.com/success",
			Failure: "https://example.com/failure",
			Pending: "https://example.com/pending",
		},
	})
	if err != nil {
		t.Fatalf("CreatePreference error: %v", err)
	}

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

func TestCheckoutProIntegration_GetPreference_Real(t *testing.T) {
	token := os.Getenv("MP_ACCESS_TOKEN")
	if token == "" {
		t.Skip("MP_ACCESS_TOKEN not set")
	}

	baseURL := os.Getenv("MP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.mercadopago.com"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     baseURL,
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	cli := mp.CheckoutPro()

	cid := checkoutpro.CurrencyIDBRL
	created, err := cli.CreatePreference(ctx, checkoutpro.CreatePreferenceInput{
		Items: []checkoutpro.PreferenceItem{
			{
				Title:      "GoPay Integration Test - GetPreference",
				Quantity:   1,
				UnitPrice:  1.00,
				CurrencyID: &cid,
			},
		},
		ExternalReference: "gopay-integration-get-preference",
	})
	if err != nil {
		t.Fatalf("CreatePreference error: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected preference id, got empty")
	}

	t.Log("Preference created")
	t.Logf("Preference ID: %s", created.ID)

	got, err := cli.GetPreference(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetPreference error: %v", err)
	}

	t.Log("Preference fetched successfully")
	t.Logf("Fetched ID: %s", got.ID)
	t.Logf("Fetched InitPoint: %s", got.InitPoint)
	t.Logf("Fetched SandboxInitPoint: %s", got.SandboxInitPoint)
	t.Logf("Fetched Items Count: %d", len(got.Items))

	if got.ID != created.ID {
		t.Fatalf("expected id %s, got %s", created.ID, got.ID)
	}

	if len(got.Items) == 0 {
		t.Fatalf("expected items, got 0")
	}
}

func TestCheckoutProIntegration_UpdatePreference_Real(t *testing.T) {
	token := os.Getenv("MP_ACCESS_TOKEN")
	if token == "" {
		t.Skip("MP_ACCESS_TOKEN not set")
	}

	baseURL := os.Getenv("MP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.mercadopago.com"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     baseURL,
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	cli := mp.CheckoutPro()
	cid := checkoutpro.CurrencyIDBRL
	created, err := cli.CreatePreference(ctx, checkoutpro.CreatePreferenceInput{
		Items: []checkoutpro.PreferenceItem{
			{
				Title:      "GoPay Integration Test - UpdatePreference",
				Quantity:   1,
				UnitPrice:  1.00,
				CurrencyID: &cid,
			},
		},
		ExternalReference: "gopay-integration-update-preference",
	})
	if err != nil {
		t.Fatalf("CreatePreference error: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected preference id, got empty")
	}

	t.Log("Preference created")
	t.Logf("Preference ID: %s", created.ID)
	newTitle := "GoPay Integration Test - Updated Title"

	updated, err := cli.UpdatePreference(ctx, created.ID, checkoutpro.UpdatePreferenceInput{
		Items: &[]checkoutpro.PreferenceItem{
			{
				Title:      newTitle,
				Quantity:   1,
				UnitPrice:  1.00,
				CurrencyID: &cid,
			},
		},
	})
	if err != nil {
		t.Fatalf("UpdatePreference error: %v", err)
	}

	t.Log("Preference created")
	t.Logf("Preference ID: %s", created.ID)
	t.Logf("InitPoint (before): %s", created.InitPoint)
	t.Logf("SandboxInitPoint (before): %s", created.SandboxInitPoint)
	if updated.ID != created.ID {
		t.Fatalf("expected id %s, got %s", created.ID, updated.ID)
	}
	if len(updated.Items) == 0 {
		t.Fatalf("expected items, got 0")
	}
	if updated.Items[0].Title != newTitle {
		t.Fatalf("expected item title %s, got %s", newTitle, updated.Items[0].Title)
	}
}
