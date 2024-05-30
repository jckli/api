package spotify

import (
	"context"
	"encoding/json"
	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
	"strconv"
)

func RecentlyPlayedHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	args := ctx.QueryArgs()
	limit, err := strconv.Atoi(string(args.Peek("limit")))
	if err != nil {
		limit = 50
	}

	redis_ctx := context.Background()

	access_token, err := redis.Do(redis_ctx, redis.B().Get().Key("spotify_access_token").Build()).
		ToString()
	if err != nil {
		ctx.Response.SetStatusCode(404)
		response := &utils.DefaultResponse{
			Status: 404,
			Data: &utils.MessageData{
				Message: "Cannot access Redis Database.",
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}
	if access_token == "" {
		refresh_data, err := getSpotifyToken()
		if err != nil {
			ctx.Response.SetStatusCode(500)
			response := &utils.DefaultResponse{
				Status: 500,
				Data: &utils.MessageData{
					Message: err.Error(),
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
		redis.Do(redis_ctx, redis.B().Set().Key("spotify_access_token").Value(refresh_data.AccessToken).Nx().Build()).
			Error()
		access_token = refresh_data.AccessToken
	}

	recently_played, err := getRecentlyPlayed(access_token, limit)
	if recently_played.Error.Status == 401 {
		refresh_data, err := getSpotifyToken()
		if err != nil {
			ctx.Response.SetStatusCode(500)
			response := &utils.DefaultResponse{
				Status: 500,
				Data: &utils.MessageData{
					Message: err.Error(),
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
		redis.Do(redis_ctx, redis.B().Set().Key("spotify_access_token").Value(refresh_data.AccessToken).Build()).
			Error()
		access_token = refresh_data.AccessToken
		recently_played, err = getRecentlyPlayed(access_token, limit)
		if err != nil {
			ctx.Response.SetStatusCode(500)
			response := &utils.DefaultResponse{
				Status: 500,
				Data: &utils.MessageData{
					Message: err.Error(),
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
	}
	if recently_played == nil {
		ctx.Response.SetStatusCode(404)
		response := &utils.DefaultResponse{
			Status: 404,
			Data: &utils.MessageData{
				Message: "No data.",
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}
	ctx.Response.SetStatusCode(200)
	response := &utils.DefaultResponse{
		Status: 200,
		Data:   recently_played,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
