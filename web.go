package framgo

import (
	"net/http"
)

// Interface to manage http request
type WebPager interface {
	// Endpoints to listen, filter function
	Endpoints() []string
	// Handler*
	Respond(vars map[string]string, r *http.Request) *HttpResponse
}
