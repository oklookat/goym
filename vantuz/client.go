package vantuz

import (
	"time"

	"golang.org/x/time/rate"
)

// HTTP Client.
type Client struct {
	// limit requests.
	limiter *rate.Limiter

	// global headers.
	headers map[string]string

	// logger.
	logger Logger

	timeout time.Duration
}

// Create request.
func (c *Client) R() *Request {
	if c.timeout <= 0 {
		c.SetTimeout(0)
	}
	return newRequest(c, c.limiter, c.timeout)
}

// Set header for all requests from this client.
func (c *Client) SetGlobalHeader(name string, value string) *Client {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[name] = value
	return c
}

// Set headers for all requests from this client.
func (c *Client) SetGlobalHeaders(h map[string]string) *Client {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	for k, v := range h {
		c.headers[k] = v
	}
	return c
}

// Set max requests per second.
//
// requests == 0 - disables limiting.
func (c *Client) SetRateLimit(requests int, per time.Duration) *Client {
	if requests == 0 || per <= 0 {
		return c
	}
	c.limiter = rate.NewLimiter(rate.Every(per), requests)
	return c
}

// Set request timeout.
//
// Default: 20 seconds.
//
// 0 and lower: set to default.
func (c *Client) SetTimeout(val time.Duration) {
	if val <= 0 {
		c.timeout = 20 * time.Second
		return
	}
	c.timeout = val
}

// Print request/response.
func (c *Client) EnableDevMode() *Client {
	c.logger.enabled = true
	return c
}

// Disable request/response print.
func (c *Client) DisableDevMode() *Client {
	c.logger.enabled = false
	return c
}
