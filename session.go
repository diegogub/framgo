package framgo

import (
	"net/http"
)

type Sessioner interface {
	Get(key string, r *http.Request) *SessionData
	Set(sd *SessionData, w http.ResponseWriter, r *http.Request)
}

type SessionData struct {
	// Expire in seconds
	Expire int64
	Data   map[string]interface{}
}

func NewSess(expire int64) *SessionData {
	var s SessionData
	s.Data = make(map[string]interface{})
	s.Expire = expire
	return &s
}
