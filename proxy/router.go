package proxy

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"sync"
)

type routerSwapper struct {
	mu     sync.Mutex
	router *mux.Router
}

func (rs *routerSwapper) Swap(newRouter *mux.Router) {
	rs.mu.Lock()
	rs.router = newRouter
	rs.mu.Unlock()
}

func (rs *routerSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.Lock()
	router := rs.router
	rs.mu.Unlock()
	router.ServeHTTP(w, r)
}

// Swapper is the main HTTP handler with the capacity to swap mux.Router.
var Swapper *routerSwapper

// LoadRouter returns a new mux.Router from services and middlewares.
func LoadRouter() *mux.Router {
	// Initiate routes.
	router := mux.NewRouter()
	for _, middleware := range middlewares {
		router.Use(middleware)
	}
	for _, service := range services {
		router.HandleFunc(service.Path, func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, service.TargetURL, http.StatusMovedPermanently)
		}).Methods(strings.Split(service.Methods, ",")...)
	}

	return router
}
