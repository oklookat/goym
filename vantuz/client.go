package vantuz

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// HTTP Client.
type Client struct {
	self *http.Client

	// limit requests.
	limiter *rate.Limiter

	// global headers.
	headers map[string]string

	log Logger
}

// Create request.
func (c *Client) R() *Request {
	return newRequest(c)
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

func (c *Client) SetLogger(log Logger) {
	c.log = log
}

func (c *Client) SetClient(cl *http.Client) {
	c.self = cl
}
