package mercadopago

import "fmt"

type Config struct {
	AccessToken string
	BaseURL     string
}

func (c Config) IsValid() error {
	if c.AccessToken == "" {
		return fmt.Errorf("%w: AccessToken is required", ErrInvalidConfig)
	}

	return nil
}
