package administration

import "net/http"

var middlewares []func(next http.Handler) http.Handler
