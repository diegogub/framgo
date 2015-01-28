package framgo

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"time"
)

type RequestContext struct {
	Method string
	AJAX   bool
	Vars   map[string]string
	URL    url.Values
	Start  time.Time
}

func NewRequestContext(r *http.Request) *RequestContext {
	var rc RequestContext
	// check if AJAX
	ajaxHead := r.Header.Get("X-Requested-With")
	if ajaxHead != "" {
		rc.AJAX = true
	}
	// HTTP Method
	rc.Method = r.Method
	rc.Vars = mux.Vars(r)
	// url form
	rc.URL = r.Form
	// request start
	rc.Start = time.Now().UTC()
	return &rc
}
