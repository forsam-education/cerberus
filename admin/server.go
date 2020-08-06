package admin

import (
	"log"
	"sync"

	"github.com/valyala/fasthttp"
)

func StartServer(group *sync.WaitGroup) {
	router := initRouter()

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
