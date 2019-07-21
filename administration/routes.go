package administration

import (
	"github.com/forsam-education/kerberos/administration/controllers"
	"github.com/forsam-education/kerberos/utils"
	"net/http"
)

var routes = []utils.Route{
	{
		Path:    "/",
		Methods: []string{http.MethodGet},
		Handler: homeHandler,
	},
	{
		Path:    "/api/services",
		Methods: []string{http.MethodGet},
		Handler: controllers.ListServices,
	},
	{
		Path:    "/api/services",
		Methods: []string{http.MethodPost},
		Handler: controllers.CreateService,
	},
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello from administration panel"))
}
