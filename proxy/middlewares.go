package proxy

import "net/http"

var middlewares = []func(next http.Handler) http.Handler{
	connectionCounter,
}

func connectionCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ConnectionCount++
		next.ServeHTTP(w, r)
		ConnectionCount--
	})
}
