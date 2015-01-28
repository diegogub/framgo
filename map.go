package framgo

import (
	"net/http"
	"sync"
)

// Helper to map request to sub function.
type Mapper struct {
	lock sync.RWMutex
	Map  map[string]func(*RequestContext, *http.Request) *HttpResponse
}

func NewMapper() *Mapper {
	var m Mapper
	m.Map = make(map[string]func(*RequestContext, *http.Request) *HttpResponse)
	return &m
}

// Register functions, to respond different actions, any HTTP action and AJAX
func (m *Mapper) Register(method string, f func(*RequestContext, *http.Request) *HttpResponse) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.Map == nil {
		m.Map = make(map[string]func(*RequestContext, *http.Request) *HttpResponse)
	}
	m.Map[method] = f
}

func (m Mapper) Respond(rc *RequestContext, r *http.Request) *HttpResponse {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	return m.lookup(rc, r)
}

func (m Mapper) lookup(rc *RequestContext, r *http.Request) *HttpResponse {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var res *HttpResponse

	fu := m.Map[rc.Method]
	if fu == nil {
		return nil
	} else {
		res = fu(rc, r)
	}
	return res
}
