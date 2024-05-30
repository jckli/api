package spotify

import (
	"context"
	"encoding/json"

	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func CurrentlyPlayingHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	market := "US"

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
		redis.Do(redis_ctx, redis.B().Set().Key("spotify_access_token").Value(refresh_data.AccessToken).Build()).
			Error()
		access_token = refresh_data.AccessToken
	}

	currently_playing, err := getCurrentlyPlaying(access_token, market)

	if currently_playing.Error.Status == 401 {
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
		currently_playing, err = getCurrentlyPlaying(access_token, market)
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
	if currently_playing == nil {
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
		Data:   currently_playing,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
