package framgo

import (
	"net/http"
	"net/url"
	"sync"
)

// Helper to map request to sub function.
type Mapper struct {
	lock sync.RWMutex
	Map  map[string]func(map[string]string, url.Values, *http.Request) *HttpResponse
}

func NewMapper() *Mapper {
	var m Mapper
	m.Map = make(map[string]func(map[string]string, url.Values, *http.Request) *HttpResponse)
	return &m
}

// Register functions, to respond different actions, any HTTP action and AJAX
func (m *Mapper) Register(method string, f func(map[string]string, url.Values, *http.Request) *HttpResponse) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.Map == nil {
		m.Map = make(map[string]func(map[string]string, url.Values, *http.Request) *HttpResponse)
	}
	m.Map[method] = f
}

func (m Mapper) Respond(vars map[string]string, u url.Values, r *http.Request) *HttpResponse {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	return m.lookup(r.Method, vars, u, r)
}

func (m Mapper) lookup(method string, vars map[string]string, u url.Values, r *http.Request) *HttpResponse {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var res *HttpResponse
	ajax := r.Header.Get("X-Requested-With")
	if ajax != "" {
		method = "AJAX"
	}

	fu := m.Map[method]
	if fu == nil {
		return nil
	} else {
		res = fu(vars, u, r)
	}
	return res
}
