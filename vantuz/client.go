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
	logger *Logger
}

// Create request.
func (c *Client) R() *Request {
	return newRequest(c, c.limiter)
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
// Example: SetRateLimit(1, 1) - 1 request per second
//
// If there are more requests, the goroutine will block.
//
// Default: 1 request per second.
func (c *Client) SetRateLimit(requests int, perSecond int) *Client {
	var dur = time.Duration(perSecond) * time.Second
	c.limiter = rate.NewLimiter(rate.Every(dur), requests)
	return c
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
