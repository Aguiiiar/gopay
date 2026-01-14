package mercadopago

import (
	"fmt"

	"github.com/Aguiiiar/gopay/providers/mercadopago/mperrors"
)

type Config struct {
	AccessToken string
	BaseURL     string
}

func (c Config) IsValid() error {
	if c.AccessToken == "" {
		return fmt.Errorf("%w: AccessToken is required", mperrors.ErrInvalidConfig)
	}

	return nil
}
