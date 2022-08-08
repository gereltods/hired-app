package app

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var route = router.New()

func StartApplication() {

	mapUrls()
	log.Fatal(fasthttp.ListenAndServe(":2999", route.Handler))
}
