package response

import (
	"net/http"
)

type JSONResponder interface {
	JSON(code int, i interface{}) error
}

type Response struct {
	responder JSONResponder
	Result    interface{} `json:"result"`
	Status    *Status     `json:"status"`
}

type Status struct {
	Code       int           `json:"code"`
	StatusText string        `json:"statusText"`
	Message    string        `json:"message,omitempty"`
	Extra      []interface{} `json:"extra,omitempty"`
}

func New(r JSONResponder, result interface{}) *Response {
	return &Response{
		responder: r,
		Result:    result,
		Status: &Status{
			Code:       http.StatusOK,
			StatusText: http.StatusText(http.StatusOK),
		},
	}
}

func (r *Response) SetStatus(code int, message string, extra ...interface{}) *Response {
	r.Status.Code = code
	r.Status.StatusText = http.StatusText(code)
	r.Status.Message = message

	if len(extra) > 0 {
		r.Status.Extra = extra
	}

	return r
}

func (r *Response) SetResult(result interface{}) *Response {
	r.Result = result
	return r
}

func (r *Response) JSON() error {
	return r.responder.JSON(http.StatusOK, r)
}
