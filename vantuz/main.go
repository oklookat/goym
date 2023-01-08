package vantuz

// Just client.
func C() *Client {
	var p = &Client{
		logger: &Logger{enabled: false},
	}
	p.SetGlobalHeaders(map[string]string{
		"User-Agent":   "oklookat/goym",
		"Content-Type": "application/json",
	}).SetRateLimit(1, 1)
	return p
}

// Client with Authorization OAuth header.
func AC(accessToken string) *Client {
	var c = C()
	c.SetGlobalHeader("Authorization", "OAuth "+accessToken)
	return c
}
