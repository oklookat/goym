// Just HTTP Client.
package vantuz

// Client.
func C() *Client {
	var p = &Client{
		logger: Logger{enabled: false},
	}
	p.SetGlobalHeader("Content-Type", "application/json").
		SetRateLimit(1, 1)
	return p
}
