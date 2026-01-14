package gopay

import "errors"

var (
	ErrNotInitialized     = errors.New("gopay not initialized: call gopay.Init(cfg) at startup")
	ErrNoProvidersEnabled = errors.New("no providers enabled")
)
