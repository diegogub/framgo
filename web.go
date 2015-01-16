package framgo

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Interface to manage http request
type WebPager interface {
	// Endpoints to listen, filter function
	Endpoints() Routes
	// Template name, or "" if plain webpage
	Template() string
	// Handler*
	Respond(vars map[string]string, r *http.Request) *HttpResponse
}

// Helper to map request to sub function
type Mapper map[string]func(map[string]string, *http.Request) *HttpResponse

func NewMapper() Mapper {
	m := make(map[string]func(map[string]string, *http.Request) *HttpResponse)
	return m
}

func (m Mapper) Respond(vars map[string]string, r *http.Request) *HttpResponse {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	return m.Map(r.URL.Path, r.Method, vars, r)
}

func (m Mapper) Register(endpoint, method string, f func(map[string]string, *http.Request) *HttpResponse) {
	m[endpoint+":"+method] = f
}

func (m Mapper) Map(endpoint, method string, vars map[string]string, r *http.Request) *HttpResponse {
	var res *HttpResponse
	fu := m[endpoint+":"+method]
	if fu == nil {
		return nil
	} else {
		res = fu(vars, r)
	}
	return res
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
