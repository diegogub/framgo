package framgo

import (
	"errors"
	"net/http"
	"sync"
)

// Interface to manage http request
type WebPager interface {
	//map endpoints to methods
	Endpoints() *RequestMap
}

type RequestMap struct {
	lock sync.Mutex
	M    map[string][]func(rc *RequestContext, r *http.Request) *HttpResponse
}

func NewRequestMap() *RequestMap {
	var rm RequestMap
	rm.M = make(map[string][]func(rc *RequestContext, r *http.Request) *HttpResponse)
	return &rm
}

func (rm *RequestMap) Add(endpoint string, f func(*RequestContext, *http.Request) *HttpResponse) error {
	rm.lock.Lock()
	defer rm.lock.Unlock()
	if endpoint == "" {
		return errors.New("Invalid endpoint")
	}

	_, ok := rm.M[endpoint]
	if !ok {
		rm.M[endpoint] = make([]func(*RequestContext, *http.Request) *HttpResponse, 0)
	}

	rm.M[endpoint] = append(rm.M[endpoint], f)

	return nil
}
