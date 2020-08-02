package admin

import (
	fasthttp "github.com/fasthttp/router"
	"github.com/forsam-education/cerberus/admin/controllers"
)

func initRouter() *fasthttp.Router {
	router := fasthttp.New()

	router.GET("/services", controllers.ListServices)
	router.POST("/services", controllers.CreateService)
	router.PUT("/services", controllers.UpdateService)
	router.PATCH("/services", controllers.UpdateService)

	return router
}
