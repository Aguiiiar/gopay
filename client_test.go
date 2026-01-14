package gopay

import "testing"

type fakeProvider struct {
	name ProviderName
}

func (p fakeProvider) Name() ProviderName {
	return p.name
}

func (p fakeProvider) ProcessPayment(input ProcessPaymentInput) (ProcessPaymentOutput, error) {
	return ProcessPaymentOutput{
		ProviderPaymentID: "fake_123",
		Status:            "paid",
	}, nil
}

func TestNew_WhenNoProviders_ReturnsErrNoProvidersEnabled(t *testing.T) {
	_, err := New(Config{})
	if err != ErrNoProvidersEnabled {
		t.Fatalf("expected ErrNoProvidersEnabled, got: %v", err)
	}
}

func TestNew_WithProvider_AllowsBuild(t *testing.T) {
	c, err := New(
		Config{},
		WithProvider(fakeProvider{name: "fake"}),
	)
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	if c == nil {
		t.Fatalf("expected client, got nil")
	}
}

func TestProvider_WhenMissing_ReturnsError(t *testing.T) {
	c, err := New(
		Config{},
		WithProvider(fakeProvider{name: "fake"}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = c.Provider("missing")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestProvider_WhenExists_ReturnsProvider(t *testing.T) {
	c, err := New(
		Config{},
		WithProvider(fakeProvider{name: "fake"}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p, err := c.Provider("fake")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Name() != "fake" {
		t.Fatalf("expected provider name fake, got: %s", p.Name())
	}
}

func TestProvider_ProcessPayment_Works(t *testing.T) {
	c, err := New(
		Config{},
		WithProvider(fakeProvider{name: "fake"}),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p, err := c.Provider("fake")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out, err := p.ProcessPayment(ProcessPaymentInput{
		Amount:   1000,
		Currency: "BRL",
		OrderID:  "order_1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ProviderPaymentID == "" {
		t.Fatalf("expected ProviderPaymentID, got empty")
	}
	if out.Status == "" {
		t.Fatalf("expected Status, got empty")
	}
}
