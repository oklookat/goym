// Just HTTP Client.
package vantuz

import (
	"net/http"
	"time"
)

const (
	_userAgent = "goym/v0.3.1 (github.com/oklookat/goym)"
)

// Client.
func C() *Client {
	p := &Client{}
	p.SetClient(&http.Client{
		Timeout: 20 * time.Second,
	})
	p.SetLogger(&dummyLogger{})

	p.SetGlobalHeader("Content-Type", "application/json")
	p.SetUserAgent(_userAgent)
	return p
}

func (c *Client) SetUserAgent(val string) {
	c.SetGlobalHeader("User-Agent", val)
}
