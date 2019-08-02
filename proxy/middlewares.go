package proxy

import (
	"github.com/forsam-education/cerberus/state"
	"github.com/gofrs/uuid"
	"net/http"
)

var middlewares = []func(next http.Handler) http.Handler{
	connectionCounter,
	injectRequestID,
}

func connectionCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := state.Manager.AddRequest()
		next.ServeHTTP(w, r)
		if err == nil {
			_ = state.Manager.RemoveRequest()
		}
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
