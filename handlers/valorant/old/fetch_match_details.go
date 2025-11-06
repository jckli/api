package valorant_old

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go/pvp"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func MatchDetailsHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	matchid := ctx.UserValue("matchid")

	auth, err := redisGetAuth(redis)
	if auth == nil || err != nil {
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

	matchDetails, err := pvp.GetMatchDetails(auth, fmt.Sprintf("%v", matchid))
	if err != nil {
		ok := auth.Reauth()
		if !ok {
			matchDetailsError(ctx)
			return
		}
		redisSetAuth(redis, auth)
		matchDetails, err = pvp.GetMatchDetails(auth, fmt.Sprintf("%v", matchid))
		if err != nil {
			matchDetailsError(ctx)
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
		Data:   matchDetails,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func matchDetailsError(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(401)
	response := &DefaultResponse{
		Status: 401,
		Data: &MessageData{
			Message: "Cannot get valorant match details.",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	return
}
