//go:build integration
// +build integration

package mercadopago_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Aguiiiar/gopay/providers/mercadopago"
	"github.com/Aguiiiar/gopay/providers/mercadopago/checkoutpro"
)

func TestMercadoPagoIntegration_CheckoutPro_Flow_Real(t *testing.T) {
	token := os.Getenv("MP_ACCESS_TOKEN")
	if token == "" {
		t.Skip("MP_ACCESS_TOKEN not set")
	}

	baseURL := os.Getenv("MP_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.mercadopago.com"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     baseURL,
		AccessToken: token,
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}
	if mp == nil {
		t.Fatalf("expected provider, got nil")
	}

	cp := mp.CheckoutPro()
	if cp == nil {
		t.Fatalf("expected checkout pro client, got nil")
	}

	cid := checkoutpro.CurrencyIDBRL
	created, err := cp.CreatePreference(ctx, checkoutpro.CreatePreferenceInput{
		Items: []checkoutpro.PreferenceItem{
			{
				Title:      "GoPay Integration - Provider Flow",
				Quantity:   1,
				UnitPrice:  1.00,
				CurrencyID: &cid,
			},
		},
		ExternalReference: "gopay-integration-provider-flow",
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
	if created.ID == "" {
		t.Fatalf("expected preference id, got empty")
	}

	t.Log("CreatePreference OK")
	t.Logf("ID: %s", created.ID)
	t.Logf("InitPoint: %s", created.InitPoint)
	t.Logf("SandboxInitPoint: %s", created.SandboxInitPoint)

	got, err := cp.GetPreference(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetPreference error: %v", err)
	}
	if got.ID != created.ID {
		t.Fatalf("expected id %s, got %s", created.ID, got.ID)
	}
	if len(got.Items) == 0 {
		t.Fatalf("expected items, got 0")
	}

	t.Log("GetPreference OK")
	t.Logf("Fetched Items: %d", len(got.Items))

	newTitle := "GoPay Integration - Provider Updated"

	updateItems := []checkoutpro.PreferenceItem{
		{
			Title:      newTitle,
			Quantity:   1,
			UnitPrice:  1.00,
			CurrencyID: &cid,
		},
	}

	updated, err := cp.UpdatePreference(ctx, created.ID, checkoutpro.UpdatePreferenceInput{
		Items: &updateItems,
	})
	if err != nil {
		t.Fatalf("UpdatePreference error: %v", err)
	}
	if updated.ID != created.ID {
		t.Fatalf("expected id %s, got %s", created.ID, updated.ID)
	}
	if len(updated.Items) == 0 {
		t.Fatalf("expected items, got 0")
	}

	t.Log("UpdatePreference OK")
	t.Logf("InitPoint (after): %s", updated.InitPoint)
	t.Logf("SandboxInitPoint (after): %s", updated.SandboxInitPoint)

	got2, err := cp.GetPreference(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetPreference after update error: %v", err)
	}
	if len(got2.Items) == 0 {
		t.Fatalf("expected items after update, got 0")
	}
	if got2.Items[0].Title != newTitle {
		t.Fatalf("expected updated title %s, got %s", newTitle, got2.Items[0].Title)
	}

	t.Log("GetPreference after Update OK")
	t.Logf("Updated Title: %s", got2.Items[0].Title)
}
