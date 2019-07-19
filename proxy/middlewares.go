package proxy

import (
	"github.com/gofrs/uuid"
	"net/http"
)

var middlewares = []func(next http.Handler) http.Handler{
	connectionCounter,
	injectRequestID,
}

func connectionCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ConnectionCount++
		next.ServeHTTP(w, r)
		ConnectionCount--
	})
}

func injectRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUUID, err := uuid.NewV4()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("X-Request-Id", requestUUID.String())
		next.ServeHTTP(w, r)
	})
}
