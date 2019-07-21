package administration

import (
	"github.com/forsam-education/kerberos/administration/controllers"
	"github.com/forsam-education/kerberos/utils"
	"net/http"
)

var apiRoutes = []utils.Route{
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
