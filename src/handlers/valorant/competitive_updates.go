package valorant

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func CompetitiveUpdatesHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	puuid := ctx.UserValue("puuid")
	args := ctx.QueryArgs()
	startIndex, err := strconv.Atoi(string(args.Peek("startIndex")))
	if err != nil {
		startIndex = 0
	}
	endIndex, err := strconv.Atoi(string(args.Peek("endIndex")))
	if err != nil {
		endIndex = 20
	}


	auth, err := redisGetAuth(redis)
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
	if auth == nil {
		authData, err := getAuth()
		if err != nil {
			ctx.Response.SetStatusCode(401)
			response := &DefaultResponse{
				Status: 401,
				Data: &MessageData{
					Message: "Cannot get valorant authentication.",
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
		redisSetAuth(redis, authData)
		auth = authData
	}
	

	competitiveUpdates, err := getCompetitiveUpdates(auth, fmt.Sprintf("%v", puuid), startIndex, endIndex)
	if err != nil {
		if err.Error() == "bad_claims" {
			reauth, err := getAuth()
			if err != nil {
				ctx.Response.SetStatusCode(401)
				response := &DefaultResponse{
					Status: 401,
					Data: &MessageData{
						Message: "Cannot get valorant authentication.",
					},
				}
				if err := json.NewEncoder(ctx).Encode(response); err != nil {
					ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
				}
				return
			}
			redisSetAuth(redis, reauth)
			competitiveUpdates, err = getCompetitiveUpdates(reauth, fmt.Sprintf("%v", puuid), startIndex, endIndex)
			if err != nil {
				ctx.Response.SetStatusCode(401)
				response := &DefaultResponse{
					Status: 401,
					Data: &MessageData{
						Message: "Cannot get valorant competitive updates.",
					},
				}
				if err := json.NewEncoder(ctx).Encode(response); err != nil {
					ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
				}
				return
			}
		} else {
			ctx.Response.SetStatusCode(401)
			response := &DefaultResponse{
				Status: 401,
				Data: &MessageData{
					Message: "Cannot get valorant competitive updates.",
				},
			}
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			}
			return
		}
	}
	if competitiveUpdates == nil {
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
		Data: competitiveUpdates,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}