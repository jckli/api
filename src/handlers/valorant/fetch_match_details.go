package valorant

import (
	"encoding/json"
	"fmt"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func MatchDetailsHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	matchid := ctx.UserValue("matchid")

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
	

	matchDetails, err := getMatchDetails(auth, fmt.Sprintf("%v", matchid))
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
			matchDetails, err = getMatchDetails(auth, fmt.Sprintf("%v", matchid))
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
	if matchDetails == nil {
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
		Data: matchDetails,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}