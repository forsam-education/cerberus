package administration

import "net/http"

var globalMiddlewares = []func(next http.Handler) http.Handler{
	setJSONHeader,
}

func setJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
