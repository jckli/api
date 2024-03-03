package valorant

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/valorant.go/pvp"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func MmrFetchPlayerHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	puuid := ctx.UserValue("puuid")

	auth, err := redisGetAuth(redis)
	if auth == nil || err != nil {
		authData, err := getAuth()
		fmt.Println(authData)
		fmt.Println(err)
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

	mmr, err := pvp.GetPlayerMmr(auth, fmt.Sprintf("%v", puuid))
	if err != nil {
		ok := auth.Reauth()
		if !ok {
			fetchPlayerMmrError(ctx)
			return
		}
		redisSetAuth(redis, auth)
		mmr, err = pvp.GetPlayerMmr(auth, fmt.Sprintf("%v", puuid))
		if err != nil {
			fetchPlayerMmrError(ctx)
			return
		}
	}
	if mmr == nil {
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
		Data:   mmr,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

func fetchPlayerMmrError(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(401)
	response := &DefaultResponse{
		Status: 401,
		Data: &MessageData{
			Message: "Cannot get valorant player mmr.",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
	return
}
