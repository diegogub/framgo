package framgo

import (
	"errors"
	"net/http"
)

// Every pager shoul
type HttpResponse struct {
	//Code
	Code int
	// Error?
	Err error
	// Redirect
	Url string
	// Cookies to set
	Cookies []*http.Cookie
	// Headers
	Headers map[string]string
	// DataResponse
	Res *Resource
}

// Generic http response
func NewHttpResponse(code int, res *Resource) *HttpResponse {
	var hr HttpResponse
	hr.Res = res
	hr.Code = code
	if code < 0 {
		hr.Code = 501
	}
	return &hr
}

func Redirect(url string, code int) *HttpResponse {
	var r HttpResponse
	r.Url = url
	r.Code = code
	return &r
}

func (ht *HttpResponse) SetResource(r *Resource) error {
	if ht.Res != nil {
		return errors.New("Resource already set!")
	}
	ht.Res = r
	return nil
}

func (ht *HttpResponse) MergeResource(datakey string, r *Resource) {
	ht.Res.Merge(datakey, r)
}
