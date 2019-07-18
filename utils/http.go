package utils

import "net/http"

// Route is the struct that represents an HTTP route to add to a router.
type Route struct {
	Path    string
	Methods []string
	Handler http.HandlerFunc
}
