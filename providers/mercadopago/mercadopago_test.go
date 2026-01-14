package mercadopago

import "testing"

func TestNew_WhenConfigInvalid_ReturnsError(t *testing.T) {
	_, err := New(Config{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestNew_WhenConfigValid_ReturnsProvider(t *testing.T) {
	p, err := New(Config{
		BaseURL:     "https://api.mercadopago.com",
		AccessToken: "token",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p == nil {
		t.Fatalf("expected provider, got nil")
	}
}
