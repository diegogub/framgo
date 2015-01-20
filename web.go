package framgo

import (
	"net/http"
	"net/url"
)

// Interface to manage http request
type WebPager interface {
	// Endpoints to listen, filter function
	Endpoints() []string
	// Http handler
	Respond(vars map[string]string, u url.Values, r *http.Request) *HttpResponse
}
