package framgo

import (
	"net/http"
	"sync"
	"time"
)

type RequestLogger interface {
	Log(rtime time.Duration, r *http.Request, res *HttpResponse)
}

// Basic logger to file and stdout
type BasicLogger struct {
	onlyRequest bool
}

type File struct {
	lock sync.RWMutex
	file string
}
