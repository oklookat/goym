// Just HTTP Client.
package vantuz

import (
	"net/http"
	"time"
)

const (
	_userAgent = "goym/v0.3.0 (github.com/oklookat/goym)"
	_prefix    = "[goym]"
)

// Client.
func C() *Client {
	p := &Client{}
	p.SetClient(&http.Client{
		Timeout: 20 * time.Second,
	})
	p.SetLogger(&dummyLogger{})

	p.SetGlobalHeader("Content-Type", "application/json")
	p.SetGlobalHeader("User-Agent", _userAgent)
	return p
}
