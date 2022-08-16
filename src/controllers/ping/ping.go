package ping

import (
	"encoding/json"
	"fasthttp/services"

	"github.com/valyala/fasthttp"
)

func Ping(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("pong")
}

func Check(ctx *fasthttp.RequestCtx) {
	//var dbs db.Dbs
	dbs := services.UsersService.GetAll()
	ctx.Response.Header.Set("Content-Type", "application/json")

	json.NewEncoder(ctx).Encode(dbs.Marshall(true))
	//ctx.Write()(db.Marshall(true))
}

func Checkwithout(ctx *fasthttp.RequestCtx) {
	str := services.UsersService.GetAllWithoutCache()
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(str)
}
