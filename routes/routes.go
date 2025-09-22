package routes

import (
	"github.com/fasthttp/router"
	"github.com/jckli/api/handlers/index"
	"github.com/jckli/api/handlers/myanimelist"
	"github.com/jckli/api/handlers/onedrive"
	"github.com/jckli/api/handlers/spotify"
	"github.com/jckli/api/handlers/valorant"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func InitRoutes(r *router.Router, redis rueidis.Client, fc *fasthttp.Client) {
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

	// onedrive routes
	r.GET("/onedrive", func(ctx *fasthttp.RequestCtx) {
		onedrive.IndexHandler(ctx, redis, fc)
	})
	r.GET("/onedrive/folder/{folderId}", func(ctx *fasthttp.RequestCtx) {
		onedrive.FolderItemsHandler(ctx, redis, fc)
	})

	// myanimelist routes
	r.GET("/myanimelist", func(ctx *fasthttp.RequestCtx) {
		mal.IndexHandler(ctx, redis, fc)
	})
	r.GET("/myanimelist/list/manga", func(ctx *fasthttp.RequestCtx) {
		mal.MangaListHandler(ctx, redis, fc)
	})
	r.GET("/myanimelist/list/anime", func(ctx *fasthttp.RequestCtx) {
		mal.AnimeListHandler(ctx, redis, fc)
	})
	r.GET("/myanimelist/list/all", func(ctx *fasthttp.RequestCtx) {
		mal.UnifiedListHandler(ctx, redis, fc)
	})
}
