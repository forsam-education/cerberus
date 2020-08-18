package administration

import (
	fasthttp "github.com/fasthttp/router"
	"github.com/forsam-education/cerberus/administration/controllers"
)

func initServicesRoutes(r *fasthttp.Group) {
	r.GET("", controllers.ListServices)
	r.POST("", controllers.CreateService)
	r.PUT("", controllers.UpdateService)
	r.PATCH("", controllers.UpdateService)
}

func initRouter() *fasthttp.Router {
	router := fasthttp.New()

	initServicesRoutes(router.Group("/services"))

	return router
}
