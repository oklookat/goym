// Just HTTP Client.
package vantuz

import (
	"context"
	"fmt"
)

const (
	errPrefix  = "goym/vantuz: "
	_userAgent = "goym/v0.2.8 (github.com/oklookat/goym)"
)

var ErrRequestCancelled = fmt.Errorf(errPrefix+"%w", context.Canceled)

// Client.
func C() *Client {
	var p = &Client{
		logger: Logger{enabled: false},
	}
	p.SetGlobalHeader("Content-Type", "application/json")
	p.SetGlobalHeader("User-Agent", _userAgent)
	return p
}
