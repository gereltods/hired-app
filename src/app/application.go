package app

import (
	"log"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var route = router.New()

func StartApplication() {
	withCors := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		//AllowedOrigins: string, // Only allow example.com to access the resource
		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: []string{"Authorization", "Accept", "Origin", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range", "x-api-role", "X-Total-Count"}, // only allow x-something-client and Content-Type in actual request
		// if you leave this empty, only simple method will be accepted
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"}, // only allow get or post to resource
		AllowCredentials: true,                                                         // resource doesn't support credentials
		AllowMaxAge:      5600,                                                         // cache the preflight result
		Debug:            true,
	})

	mapUrls()
	if err := fasthttp.ListenAndServe(":2999", withCors.CorsMiddleware(route.Handler)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
