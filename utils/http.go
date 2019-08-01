package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ResponseExtraData is the type for response specific data.
type ResponseExtraData map[string]interface{}

const (
	// HeaderXForwardedFor is an HTTP header.
	HeaderXForwardedFor = "X-Forwarded-For"
	// HeaderContentType is an HTTP header.
	HeaderContentType   = "Content-Type"
)

// Route is the struct that represents an HTTP route to add to a router.
type Route struct {
	Path        string
	Methods     []string
	Handler     http.HandlerFunc
	Middlewares []mux.MiddlewareFunc
}

// Response is the basic response type for JSON responses.
type Response struct {
	Status string            `json:"status"`
	Data   ResponseExtraData `json:"data"`
}
