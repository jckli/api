package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/rueian/rueidis"
	"github.com/jckli/api/src/handlers/index"
	"github.com/jckli/api/src/handlers/spotify"
)

func InitRoutes(r *router.Router, redis rueidis.Client) {
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		index.IndexHandler(ctx, redis)
	})
	r.GET("/spotify/top-items/{itype}", func(ctx *fasthttp.RequestCtx) {
		spotify.TopItemsHandler(ctx, redis)
	})
	r.GET("/spotify/currently-playing", func(ctx *fasthttp.RequestCtx) {
		spotify.CurrentlyPlayingHandler(ctx, redis)
	})
}
