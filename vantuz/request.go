package vantuz

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func newRequest(cl *Client) *Request {
	r := &Request{
		cl:      cl,
		headers: make(map[string]string),
	}

	if len(cl.headers) == 0 {
		return r
	}

	for k, v := range cl.headers {
		r.headers[k] = v
	}

	return r
}

// HTTP Request.
type Request struct {
	cl *Client

	// query params.
	params url.Values

	bodyStr string

	// request headers.
	headers map[string]string

	// unmarshal body (HTTP error)
	err any

	// unmarshal body (HTTP success)
	result any
}

// GET request.
func (r Request) Get(ctx context.Context, url string) (*Response, error) {
	resp, err := r.exec(ctx, http.MethodGet, url)
	if err != nil {
		r.cl.log.Err("", err)
	}
	return resp, err
}

// POST request.
func (r Request) Post(ctx context.Context, url string) (*Response, error) {
	resp, err := r.exec(ctx, http.MethodPost, url)
	if err != nil {
		r.cl.log.Err("", err)
	}
	return resp, err
}

// DELETE request.
func (r Request) Delete(ctx context.Context, url string) (*Response, error) {
	resp, err := r.exec(ctx, http.MethodDelete, url)
	if err != nil {
		r.cl.log.Err("", err)
	}
	return resp, err
}

// Unmarshall body in 'err' if status code >= 400.
//
// Does nothing if param 'err' be nil.
func (r *Request) SetError(err any) *Request {
	if err == nil {
		return r
	}
	r.err = err
	return r
}

// Unmarshall body if status code >= 200 and <= 299.
//
// Does nothing if param 'res' be nil.
func (r *Request) SetResult(res any) *Request {
	if res == nil {
		return r
	}
	r.result = res
	return r
}

func (r *Request) setStringBody(val string, contentType string) {
	r.bodyStr = val
	r.setContentType(contentType)
	r.setContentLength(len(val))
}

// application/x-www-form-urlencoded
func (r *Request) SetFormUrlValues(data url.Values) *Request {
	if len(data) == 0 {
		return r
	}
	encoded := data.Encode()
	r.setStringBody(encoded, "application/x-www-form-urlencoded")
	return r
}

// application/x-www-form-urlencoded
func (r *Request) SetFormUrlMap(data map[string]string) *Request {
	if len(data) == 0 {
		return r
	}

	vals := url.Values{}
	for k, v := range data {
		vals.Set(k, v)
	}

	encoded := vals.Encode()
	r.setStringBody(encoded, "application/x-www-form-urlencoded")

	return r
}

// application/json
func (r *Request) SetJsonString(data string) *Request {
	r.setStringBody(data, "application/json")
	return r
}

// Set query params.
func (r *Request) SetQueryParams(params url.Values) *Request {
	r.params = params
	return r
}

func (r *Request) setContentType(val string) *Request {
	r.headers["Content-Type"] = val
	return r
}

func (r *Request) setContentLength(val int) {
	r.headers["Content-Length"] = strconv.Itoa(val)
}

// Execute request.
func (r *Request) exec(ctx context.Context, method string, urld string) (*Response, error) {
	r.cl.log.Debugf("%s: %s", method, urld)

	body := strings.NewReader(r.bodyStr)

	// validate url.
	if _, err := url.Parse(urld); err != nil {
		return nil, err
	}

	// create request.
	req, err := http.NewRequestWithContext(ctx, method, urld, body)
	if err != nil {
		return nil, err
	}

	// set headers.
	if len(r.headers) > 0 {
		for k, v := range r.headers {
			r.cl.log.Debugf(`set header: "%s": "%s"`, k, v)
			req.Header.Set(k, v)
		}
	}

	if len(r.params) > 0 {
		req.URL.RawQuery = r.params.Encode()
		r.cl.log.Debugf("query: %s", req.URL.RawQuery)
	}

	// wait limiter
	if r.cl.limiter != nil {
		r.cl.log.Debugf("limiter wait...")
		if err = r.cl.limiter.Wait(ctx); err != nil {
			return nil, err
		}
	}

	// send request.
	hResp, err := r.cl.self.Do(req)
	if err != nil {
		return nil, err
	}

	// make response.
	if err = r.unmarshalResponse(hResp); err != nil {
		return nil, err
	}

	return newResponse(r, hResp), err
}

// Unmarshal response body to result/err.
func (r Request) unmarshalResponse(resp *http.Response) error {
	if resp.Body == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if r.err != nil && isHttpError(resp.StatusCode) {
		return json.Unmarshal(body, r.err)
	}

	if r.result != nil && isHttpSuccess(resp.StatusCode) {
		return json.Unmarshal(body, r.result)
	}

	return err
}
