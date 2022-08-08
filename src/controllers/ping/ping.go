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
	ctx.Write(str)
}

func Checkwithout(ctx *fasthttp.RequestCtx) {
	str := services.UsersService.GetAllWithoutCache()

	ctx.Write(str)
}
