package administration

import (
	"github.com/forsam-education/cerberus/administration/controllers"
	"github.com/forsam-education/cerberus/utils"
	"net/http"
)

var apiRoutes = []utils.Route{
	{
		Path:    "/services",
		Methods: []string{http.MethodGet},
		Handler: controllers.ListServices,
	},
	{
		Path:    "/services",
		Methods: []string{http.MethodPost},
		Handler: controllers.CreateService,
	},
	{
		Path:    "/services",
		Methods: []string{http.MethodPatch, http.MethodPut},
		Handler: controllers.UpdateService,
	},
	{
		Path:    "/status",
		Methods: []string{http.MethodGet},
		Handler: controllers.Status,
	},
}
