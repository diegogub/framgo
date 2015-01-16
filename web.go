package framgo

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Interface to manage http request
type WebPager interface {
	// Endpoints to listen, filter function
	Endpoints() Routes
	// Template name
	Template() string
	// Handler
	Respond(vars map[string]string, r *http.Request) *HttpResponse
	// Type: html,plain
	Type() string
}

// Routes represent and enpoint and corresponding filter function
type Routes map[string]func(*mux.Route)

func (r Routes) AddRoute(endpoint string, filterFunc func(*mux.Route)) {
	r[endpoint] = filterFunc
}

func NewRoutes() Routes {
	r := make(map[string]func(*mux.Route))
	return r
}
