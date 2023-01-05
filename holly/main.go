package holly

import "github.com/imroc/req/v3"

/**
holly is a wrapper over 3rd party HTTP client.

This project uses an HTTP client,
and being tied to a specific one is not very good.

If you need to replace the client, this can be done here,
without breaking other code.
*/

// Just client.
func C() *Client {
	var p = &Client{}
	p.new()
	return p
}

// Client with Authorization header.
func AC(accessToken string) *Client {
	var c = C()
	c.SetGlobalHeader("Authorization", "OAuth "+accessToken)
	return c
}

// Wrapper.
type Client struct {
	low *req.Client
}

func (c *Client) new() {
	c.low = req.C()
	c.low.SetUserAgent("oklookat/goym").SetCommonHeader("Content-Type", "application/json")
}

func (c *Client) EnableDevMode() {
	c.low.EnableDumpAll()
	c.low.EnableDebugLog()
	c.low.EnableTraceAll()
}

func (c *Client) DisableDevMode() {
	c.low.DisableDumpAll()
	c.low.DisableDebugLog()
	c.low.DisableTraceAll()
}

// Create request.
func (c *Client) R() *Request {
	var request = &Request{}
	request.new(c.low)
	return request
}

// Set header for all requests from this client.
func (c *Client) SetGlobalHeader(name, value string) *Client {
	c.low.SetCommonHeader(name, value)
	return c
}

// Set headers for all requests from this client.
func (c *Client) SetGlobalHeaders(h map[string]string) *Client {
	c.low.SetCommonHeaders(h)
	return c
}
