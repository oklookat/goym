package holly

import (
	"github.com/imroc/req/v3"
)

type Request struct {
	low *req.Request
}

func (r *Request) new(c *req.Client) {
	r.low = c.R()
}

// Set request body.
//
// allowed: string, []byte, io.Reader, map, struct.
//
// Does nothing if param 'body' be nil.
func (r *Request) SetBody(body interface{}) *Request {
	if body == nil {
		return r
	}
	r.low.SetBody(body)
	return r
}

// Unmarshall body in 'err' if status code >= 400.
//
// Does nothing if param 'err' be nil.
func (r *Request) SetError(err interface{}) *Request {
	if err == nil {
		return r
	}
	r.low.SetError(err)
	return r
}

// Unmarshall body if status code >= 200 and <= 299.
//
// Does nothing if param 'res' be nil.
func (r *Request) SetResult(res interface{}) *Request {
	if res == nil {
		return r
	}
	r.low.SetResult(res)
	return r
}

// SetFormData set the form data from a map, will not been used if request method does not allow payload.
//
// Does nothing if param 'data' be nil.
func (r *Request) SetFormData(data map[string]string) *Request {
	if data == nil {
		return r
	}
	r.low.SetFormData(data)
	return r
}

// AddQueryParam add a URL query parameter for the request.
func (r *Request) AddQueryParam(key string, value string) *Request {
	r.low.AddQueryParam(key, value)
	return r
}

// SetQueryParams set URL query parameters from a map for the request.
func (r *Request) SetQueryParams(params map[string]string) *Request {
	r.low.SetQueryParams(params)
	return r
}

// GET request.
func (r *Request) Get(url string) (resp *Response, err error) {
	resp = &Response{}

	var lowResp *req.Response
	lowResp, err = r.low.Get(url)
	resp.new(lowResp)
	if err != nil {
		return
	}

	return
}

// POST request.
func (r *Request) Post(url string) (resp *Response, err error) {
	resp = &Response{}

	var lowResp *req.Response
	lowResp, err = r.low.Post(url)
	resp.new(lowResp)
	if err != nil {
		return
	}

	return
}

// DELETE request.
func (r *Request) Delete(url string) (resp *Response, err error) {
	resp = &Response{}

	var lowResp *req.Response
	lowResp, err = r.low.Delete(url)
	resp.new(lowResp)
	if err != nil {
		return
	}

	return
}
