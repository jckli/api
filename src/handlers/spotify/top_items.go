package spotify

import (
	"encoding/json"
	"context"
	"strconv"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/rueian/rueidis"
)

func TopItemsHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	item_type := ctx.UserValue("itype")
	args := ctx.QueryArgs()
	timeRange := string(args.Peek("time_range"))
	if timeRange == "" {
		timeRange = "long_term"
	}
	limit, err := strconv.Atoi(string(args.Peek("limit")))
	if err != nil {
		limit = 50
	}

	redis_ctx := context.Background()

	access_token, err := redis.Do(redis_ctx, redis.B().Get().Key("spotify_access_token").Build()).ToString()
	if err != nil {
		ctx.Response.SetStatusCode(401)
		response := &DefaultResponse{
			Status: 401,
			Data: &MessageData{
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
			ctx.Response.SetStatusCode(401)
			response := &DefaultResponse{
				Status: 401,
				Data: &MessageData{
					Message: "Cannot refresh Spotify token.",
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
		redis.Do(redis_ctx, redis.B().Set().Key("spotify_access_token").Value(refresh_data.AccessToken).Nx().Build()).Error()
		access_token = refresh_data.AccessToken
	}

	fmt.Println(access_token)

	top_items, err := getTopItems(access_token, fmt.Sprintf("%v", item_type), timeRange, limit)
	if err == fmt.Errorf("Unauthorized") {
		refresh_data, err := getSpotifyToken()
		if err != nil {
			ctx.Response.SetStatusCode(401)
			response := &DefaultResponse{
				Status: 401,
				Data: &MessageData{
					Message: "Cannot refresh Spotify token.",
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
		redis.Do(redis_ctx, redis.B().Set().Key("spotify_access_token").Value(refresh_data.AccessToken).Nx().Build()).Error()
		access_token = refresh_data.AccessToken
		top_items, err = getTopItems(access_token, fmt.Sprintf("%v", item_type), timeRange, limit)
		if err != nil {
			ctx.Response.SetStatusCode(401)
			response := &DefaultResponse{
				Status: 401,
				Data: &MessageData{
					Message: "Cannot get Spotify top items.",
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
	} else if err != nil {
		ctx.Response.SetStatusCode(401)
		response := &DefaultResponse{
			Status: 401,
			Data: &MessageData{
				Message: "Cannot get Spotify top items.",
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}
	if top_items == nil {
		ctx.Response.SetStatusCode(404)
		response := &DefaultResponse{
			Status: 404,
			Data: &MessageData{
				Message: "No data.",
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}
	ctx.Response.SetStatusCode(200)
	response := &DefaultResponse{
		Status: 200,
		Data: top_items,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}