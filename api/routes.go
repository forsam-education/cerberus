package api

import (
	"github.com/forsam-education/kerberos/utils"
	"net/http"
)

func getRoutes() []utils.Route {
	return []utils.Route{
		{
			Path:    "/",
			Methods: []string{utils.GET},
			Handler: homeHandler,
		},
	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello from administration panel"))
}
