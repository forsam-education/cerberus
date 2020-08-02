package admin

import (
	"log"

	"github.com/valyala/fasthttp"
)

func StartServer() {
	router := initRouter()

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
