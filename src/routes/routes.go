package routes

import (
	"github.com/fasthttp/router"
	"github.com/jckli/api/src/handlers/index"
	"github.com/jckli/api/src/handlers/spotify"
	"github.com/jckli/api/src/handlers/valorant"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func InitRoutes(r *router.Router, redis rueidis.Client) {
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		index.IndexHandler(ctx, redis)
	})

	// spotify routes
	r.GET("/spotify", func(ctx *fasthttp.RequestCtx) {
		spotify.IndexHandler(ctx, redis)
	})
	r.GET("/spotify/top-items/{itype}", func(ctx *fasthttp.RequestCtx) {
		spotify.TopItemsHandler(ctx, redis)
	})
	r.GET("/spotify/currently-playing", func(ctx *fasthttp.RequestCtx) {
		spotify.CurrentlyPlayingHandler(ctx, redis)
	})
	r.GET("/spotify/recently-played", func(ctx *fasthttp.RequestCtx) {
		spotify.RecentlyPlayedHandler(ctx, redis)
	})

	// valorant routes
	r.GET("/valorant", func(ctx *fasthttp.RequestCtx) {
		valorant.IndexHandler(ctx, redis)
	})
	r.GET("/valorant/mmr/players/{puuid}", func(ctx *fasthttp.RequestCtx) {
		valorant.MmrFetchPlayerHandler(ctx, redis)
	})
	r.GET("/valorant/mmr/players/{puuid}/competitive-updates", func(ctx *fasthttp.RequestCtx) {
		valorant.CompetitiveUpdatesHandler(ctx, redis)
	})
	r.GET("/valorant/match-details/{matchid}", func(ctx *fasthttp.RequestCtx) {
		valorant.MatchDetailsHandler(ctx, redis)
	})
}
