// Just HTTP Client.
package vantuz

import "errors"

const (
	errPrefix = "goym/vantuz: "
)

var (
	ErrRequestCancelled = errors.New(errPrefix + "request cancelled")
	ErrNilRequestBefore = errors.New(errPrefix + "nil Request.before()")
	ErrResponse         = errors.New(errPrefix + "nil http.Response")
)

// Client.
func C() *Client {
	var p = &Client{
		logger: Logger{enabled: false},
	}
	p.SetGlobalHeader("Content-Type", "application/json").
		SetRateLimit(1, 1)
	return p
}
