package app

import (
	"fasthttp/controllers/ping"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)
	route.GET("/check", ping.Check)
	route.GET("/checkwithout", ping.Checkwithout)
}
