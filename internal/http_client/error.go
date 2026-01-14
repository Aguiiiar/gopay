package httpclient

import "errors"

var (
	ErrBadStatus = errors.New("httpclient: bad status code")
)
