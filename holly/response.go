package holly

import (
	"github.com/imroc/req/v3"
)

type Response struct {
	low *req.Response
}

func (r *Response) new(resp *req.Response) {
	r.low = resp
}

// status code >= 400.
func (r *Response) IsError() bool {
	return r.low.IsError()
}

// status code >= 200 and <= 299.
func (r *Response) IsSuccess() bool {
	return r.low.IsSuccess()
}

// Object from Request.SetError().
func (r *Response) Error() any {
	return r.low.Error()
}
