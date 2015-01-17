package framgo

import (
	"errors"
	"net/http"
)

const HTML_TYPE = "html"
const PLAIN_TYPE = "plain"

// Every pager shoul
type HttpResponse struct {
	//Type : html or plain. Html use WebPager template and resource Data, while Plain just the Plain resource
	Type string
	// Template name
	Template string
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
func NewHttpResponse(code int, res *Resource, template string, resType string) *HttpResponse {
	var hr HttpResponse
	switch resType {
	case "html", "plain":
		hr.Type = resType
	default:
		// default plain response
		hr.Type = PLAIN_TYPE
	}

	hr.Template = template
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

func (ht *HttpResponse) SetCookie(c *http.Cookie) {
	if ht.Cookies == nil {
		ht.Cookies = make([]*http.Cookie, 0)
	}
	ht.Cookies = append(ht.Cookies, c)
}

func (ht *HttpResponse) SetHeader(key, value string) {
	if ht.Headers == nil {
		ht.Headers = make(map[string]string)
	}
	ht.Headers[key] = value
}

func (ht *HttpResponse) MergeResource(datakey string, r *Resource) {
	ht.Res.Merge(datakey, r)
}
