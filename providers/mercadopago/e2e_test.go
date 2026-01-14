package mercadopago_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aguiiiar/gopay/providers/mercadopago"
)

func TestE2E_CreatePreference(t *testing.T) {
	t.Parallel()

	type reqBody struct {
		Items []struct {
			Title      string  `json:"title"`
			Quantity   int     `json:"quantity"`
			CurrencyID string  `json:"currency_id"`
			UnitPrice  float64 `json:"unit_price"`
		} `json:"items"`
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1) valida método e path
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/checkout/preferences" {
			t.Fatalf("expected /checkout/preferences, got %s", r.URL.Path)
		}

		// 2) valida auth
		if got := r.Header.Get("Authorization"); got != "Bearer token_123" {
			t.Fatalf("expected Authorization 'Bearer token_123', got %q", got)
		}

		// 3) valida JSON enviado
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read body error: %v", err)
		}
		defer r.Body.Close()

		var body reqBody
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Fatalf("invalid json: %v. raw=%s", err, string(raw))
		}

		if len(body.Items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(body.Items))
		}
		if body.Items[0].Title != "Produto A" {
			t.Fatalf("expected title 'Produto A', got %q", body.Items[0].Title)
		}
		if body.Items[0].Quantity != 2 {
			t.Fatalf("expected quantity 2, got %d", body.Items[0].Quantity)
		}
		if body.Items[0].CurrencyID != "BRL" {
			t.Fatalf("expected currency BRL, got %q", body.Items[0].CurrencyID)
		}

		// 4) responde como Mercado Pago faria
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{
			"id":"pref_1",
			"init_point":"https://pay.example/init",
			"sandbox_init_point":"https://pay.example/sandbox"
		}`))
	}))
	defer srv.Close()

	// provider real + httpclient interno
	mp, err := mercadopago.New(mercadopago.Config{
		BaseURL:     srv.URL,
		AccessToken: "token_123",
	})
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	out, err := mp.CreatePreference(context.Background(), mercadopago.CreatePreferenceInput{
		Items: []mercadopago.PreferenceItem{
			{
				Title:      "Produto A",
				Quantity:   2,
				CurrencyID: "BRL",
				UnitPrice:  10.50,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreatePreference error: %v", err)
	}

	if out.ID != "pref_1" {
		t.Fatalf("expected id pref_1, got %q", out.ID)
	}
	if out.InitPoint == "" {
		t.Fatalf("expected init_point, got empty")
	}
}
