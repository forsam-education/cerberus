package administration

import (
	"github.com/forsam-education/kerberos/utils"
	"net/http"
)

var routes = []utils.Route{
	{
		Path:    "/",
		Methods: []string{http.MethodGet},
		Handler: homeHandler,
	},
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello from administration panel"))
}
