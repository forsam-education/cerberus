package administration

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"net/http"
)

var globalMiddlewares = []func(next http.Handler) http.Handler{
	setJSONHeader,
	verboseRouteLogger,
}

func setJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func verboseRouteLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogVerbose(fmt.Sprintf("Matched API route %s from %s", r.URL.Path, r.RemoteAddr), nil)
		next.ServeHTTP(w, r)
	})
}
