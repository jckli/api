package valorant

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go/pvp"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
	"strconv"
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

	competitiveUpdates, err := pvp.GetCompetitiveUpdates(auth, fmt.Sprintf("%v", puuid), pvp.WithStartIndex(startIndex), pvp.WithEndIndex(endIndex))
	if err != nil {
		ok := auth.Reauth()
		if !ok {
			competitiveUpdatesError(ctx)
			return
		}
		redisSetAuth(redis, auth)
		competitiveUpdates, err = pvp.GetCompetitiveUpdates(auth, fmt.Sprintf("%v", puuid), pvp.WithStartIndex(startIndex), pvp.WithEndIndex(endIndex))
		if err != nil {
			competitiveUpdatesError(ctx)
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
		Data:   competitiveUpdates,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func competitiveUpdatesError(ctx *fasthttp.RequestCtx) {
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
