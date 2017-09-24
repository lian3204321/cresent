package context

import (
	"net/http"
	"io"
)

// 自定义响应
type Response struct {
	resp   *http.ResponseWriter
	status int
	writer io.Writer
}

func NewResponse(w *http.ResponseWriter) *Response {
	r := new(Response)
	r.resp = w
	return r
}
