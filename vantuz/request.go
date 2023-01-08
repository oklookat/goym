package vantuz

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/time/rate"
)

func newRequest(cl *Client, limit *rate.Limiter) *Request {
	return &Request{
		cl:      cl,
		limiter: limit,
		headers: make(map[string]string),
	}
}

// HTTP Request.
type Request struct {
	cl      *Client
	limiter *rate.Limiter

	// query params.
	params url.Values

	body io.Reader

	// request headers.
	headers map[string]string

	// unmarshal body (HTTP error)
	err any

	// unmarshal body (HTTP success)
	result any
}

// GET request.
func (r *Request) Get(url string) (*Response, error) {
	return r.exec(http.MethodGet, url)
}

// POST request.
func (r *Request) Post(url string) (*Response, error) {
	return r.exec(http.MethodPost, url)
}

// DELETE request.
func (r *Request) Delete(url string) (*Response, error) {
	return r.exec(http.MethodDelete, url)
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

// application/x-www-form-urlencoded
func (r *Request) SetFormData(data map[string]string) *Request {
	if data == nil {
		return r
	}

	var vals = url.Values{}
	for k, v := range data {
		vals.Set(k, v)
	}

	var encoded = vals.Encode()
	r.body = strings.NewReader(encoded)
	r.setContentType("application/x-www-form-urlencoded")
	r.setContentLength(len(encoded))

	return r
}

// Set request query params.
func (r *Request) SetQueryParams(params map[string]string) *Request {
	if params == nil {
		return r
	}
	if r.params == nil {
		r.params = make(url.Values)
	}

	for k, v := range params {
		r.params.Set(k, v)
	}

	return r
}

func (r *Request) setContentType(val string) *Request {
	r.headers["Content-Type"] = val
	return r
}

func (r *Request) setContentLength(val int) {
	r.headers["Content-Length"] = strconv.Itoa(val)
}

// Before request.
func (r *Request) before(req *http.Request) error {
	if req == nil {
		return errors.New("nil before()")
	}

	for k, v := range r.cl.headers {
		req.Header.Set(k, v)
	}
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	req.URL.RawQuery = r.params.Encode()
	ctx := context.Background()
	err := r.limiter.Wait(ctx)

	return err
}

// Execute request.
func (r *Request) exec(method string, urld string) (resp *Response, err error) {
	defer func() {
		if err != nil {
			r.cl.logger.log(err.Error())
		}
	}()

	if _, err := url.Parse(urld); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urld, r.body)
	if err != nil {
		return nil, err
	}

	if err = r.before(req); err != nil {
		return nil, err
	}

	go func() {
		r.cl.logger.log("==== Request: %v ====", req.URL)
		for k, v := range req.Header {
			r.cl.logger.log("%s: %s", k, strings.Join(v, ","))
		}
	}()

	client := &http.Client{}
	hResp, err := client.Do(req)
	r.cl.logger.log("==== Response: %v (%v) ====", req.URL, hResp.StatusCode)
	if err != nil {
		return
	}

	if err = r.unmarshalResponse(hResp); err != nil {
		return
	}

	return newResponse(r, hResp), err
}

// Unmarshal response body to result/err.
func (r *Request) unmarshalResponse(resp *http.Response) error {
	if resp == nil {
		return errors.New("nil resp")
	}
	if resp.Body == nil {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r.cl.logger.log("Body: %s", string(body))

	if r.err != nil && isHttpError(resp.StatusCode) {
		return json.Unmarshal(body, r.err)
	}

	if r.result != nil && isHttpSuccess(resp.StatusCode) {
		return json.Unmarshal(body, r.result)
	}

	return err
}
