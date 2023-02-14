package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/jckli/api/src/handlers/index"
	"github.com/rueian/rueidis"
)

func InitRoutes(r *router.Router, redis rueidis.Client) {
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		index.IndexHandler(ctx, redis)
	})
}
