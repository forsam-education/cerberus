package utils

import "net/http"

// Route is the struct that represents an HTTP route to add to a router.
type Route struct {
	Path    string
	Methods []string
	Handler http.HandlerFunc
}

const (
	// GET HTTP METHOD
	GET string = "GET"
	// POST HTTP METHOD
	POST string = "POST"
	// PUT HTTP METHOD
	PUT string = "PUT"
	// PATCH HTTP METHOD
	PATCH string = "PATCH"
	// HEAD HTTP METHOD
	HEAD string = "HEAD"
)
