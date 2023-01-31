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
	"time"

	"golang.org/x/time/rate"
)

func newRequest(cl *Client, limit *rate.Limiter, timeout time.Duration) *Request {
	var r = &Request{
		cl:      cl,
		limiter: limit,
		headers: make(map[string]string),
		timeout: timeout,
	}
	if cl.headers != nil {
		for k, v := range cl.headers {
			r.headers[k] = v
		}
	}
	return r
}

// HTTP Request.
type Request struct {
	cl      *Client
	limiter *rate.Limiter

	// query params.
	params url.Values

	bodyStr string

	// request headers.
	headers map[string]string

	// unmarshal body (HTTP error)
	err any

	// unmarshal body (HTTP success)
	result any

	timeout time.Duration
}

// GET request.
func (r Request) Get(ctx context.Context, url string) (*Response, error) {
	return r.exec(ctx, http.MethodGet, url)
}

// POST request.
func (r Request) Post(ctx context.Context, url string) (*Response, error) {
	return r.exec(ctx, http.MethodPost, url)
}

// DELETE request.
func (r Request) Delete(ctx context.Context, url string) (*Response, error) {
	return r.exec(ctx, http.MethodDelete, url)
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
	if data == nil {
		return r
	}
	var encoded = data.Encode()
	r.setStringBody(encoded, "application/x-www-form-urlencoded")
	return r
}

// application/x-www-form-urlencoded
func (r *Request) SetFormUrlMap(data map[string]string) *Request {
	if len(data) == 0 {
		return r
	}
	var vals = url.Values{}
	for k, v := range data {
		vals.Set(k, v)
	}
	var encoded = vals.Encode()
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
func (r *Request) exec(ctx context.Context, method string, urld string) (resp *Response, err error) {
	defer func() {
		if err != nil {
			if errors.Is(err, context.Canceled) {
				err = ErrRequestCancelled
			}
			r.cl.logger.log(err.Error())
		}
	}()

	var body = strings.NewReader(r.bodyStr)

	// validate url.
	if _, err := url.Parse(urld); err != nil {
		return nil, err
	}

	// create request.
	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, method, urld, body)
	if err != nil {
		return nil, err
	}

	// set headers.
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
	req.URL.RawQuery = r.params.Encode()
	if r.limiter != nil {
		if err = r.limiter.Wait(ctx); err != nil {
			return
		}
	}

	r.cl.logger.request(req)

	client := &http.Client{
		Timeout: r.timeout,
	}

	// send request.
	var hResp *http.Response
	hResp, err = client.Do(req)
	if err != nil {
		return
	}

	// make response.
	r.cl.logger.response(hResp)
	if err = r.unmarshalResponse(hResp); err != nil {
		return
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
	r.cl.logger.responseBody(body)

	if r.err != nil && isHttpError(resp.StatusCode) {
		return json.Unmarshal(body, r.err)
	}

	if r.result != nil && isHttpSuccess(resp.StatusCode) {
		return json.Unmarshal(body, r.result)
	}

	return err
}
