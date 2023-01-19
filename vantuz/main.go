// Just HTTP Client.
package vantuz

import (
	"context"
	"fmt"
)

const (
	errPrefix = "goym/vantuz: "
)

var ErrRequestCancelled = fmt.Errorf(errPrefix+"%w", context.Canceled)

// Client.
func C() *Client {
	var p = &Client{
		logger: Logger{enabled: false},
	}
	p.SetGlobalHeader("Content-Type", "application/json").
		SetRateLimit(1, 1)
	return p
}
