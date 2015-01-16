package main

import (
	"net/http"
)

// Interface to manage http request
type WebPager interface {
	// Register Methods
	// Endpoints to listen
	Endpoints() []string
	// Returns http method to respond
	Method() []string
	// headers filter
	Headers() map[string]string
	// Template name
	Template() string
	// Handler
	Respond(vars map[string]string, r *http.Request) *HttpResponse
	// Type: html,plain
	Type() string
}
