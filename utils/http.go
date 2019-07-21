package utils

import (
	"github.com/gorilla/mux"
	"net/http"
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
	Status string `json:"status"`
}
