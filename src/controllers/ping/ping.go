package ping

import (
	"fasthttp/services"

	"github.com/valyala/fasthttp"
)

func Ping(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("pong")
}

func Check(ctx *fasthttp.RequestCtx) {
	str := services.UsersService.GetAll()
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(str)
}

func Checkwithout(ctx *fasthttp.RequestCtx) {
	str := services.UsersService.GetAllWithoutCache()
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(str)
}
