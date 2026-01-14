# GoPay

GoPay is a Go library that helps you integrate payment gateways in a simple and consistent way.

The project aims to abstract payment providers behind clear, explicit APIs, so you can focus on your business logic instead of provider-specific details.

GoPay is designed as a **library**, not a service or framework.

> ⚠️ This project is under active development and is not ready for production use yet.

---

## Project Status

**Development stage:** Early / In progress

The public API may change until the first stable release.

---

## Goals

• Simplify payment gateway integrations  
• Provide explicit and well-typed APIs  
• Avoid hidden magic or global state  
• Keep provider logic isolated and organized  
• Support real-world payment flows  

---

## Current Provider

### Mercado Pago

Status: **In progress**

Currently being developed:

• Checkout Pro  
• Create Preference  
• Get Preference  
• Typed error handling  
• Real integration tests  

Planned next steps:

• Checkout API (Payments)  
• Pix  
• Boleto  
• Webhooks  
• Payment status normalization  

---

## Example (Checkout Pro)

```go
mp, err := mercadopago.New(mercadopago.Config{
	AccessToken: "<YOUR_ACCESS_TOKEN>",
})
if err != nil {
	log.Fatal(err)
}

out, err := mp.CheckoutPro().CreatePreference(ctx, checkoutpro.CreatePreferenceInput{
	Items: []checkoutpro.PreferenceItem{
		{
			Title:     "My Product",
			Quantity:  1,
			UnitPrice: 10.00,
		},
	},
})
if err != nil {
	log.Fatal(err)
}

fmt.Println(out.InitPoint)
```

---

## Roadmap

• Finalize Mercado Pago provider  
• Add Checkout API support  
• Add Pix and Boleto flows  
• Add Webhook handling  
• Define provider-agnostic interfaces  
• Add new providers:
  - PayPal  
  - Pagar.me  
  - Efí Pay  
  - Asaas  

---

## Design Principles

• Explicit over implicit  
• One provider, one responsibility  
• Clear separation of payment flows  
• Go-first design  
• No framework lock-in  

---

## Development

Run tests:

```bash
go test ./...
```

Run integration tests:

```bash
go test ./... -tags=integration -v
```

---