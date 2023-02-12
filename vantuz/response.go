package vantuz

import (
	"net/http"
)

func newResponse(req *Request, resp *http.Response) *Response {
	return &Response{
		req:  req,
		self: resp,
	}
}

type Response struct {
	self *http.Response
	req  *Request
}

// status code >= 400.
func (r Response) IsError() bool {
	return isHttpError(r.self.StatusCode)
}

// status code >= 200 and <= 299.
func (r Response) IsSuccess() bool {
	return isHttpSuccess(r.self.StatusCode)
}

// Object from Request.SetError().
func (r Response) Error() any {
	return r.req.err
}
