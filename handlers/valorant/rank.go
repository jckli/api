package valorant

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func RankHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client, client *fasthttp.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))

	puuid := os.Getenv("VALORANT_PUUID")
	if puuid == "" {
		err := fmt.Errorf("VALORANT_PUUID environment variable is not set")
		fmt.Printf("Error: %v\n", err)
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		response := &utils.DefaultResponse{
			Status: fasthttp.StatusInternalServerError,
			Data: &utils.MessageData{
				Message: err.Error(),
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}

	rankData, err := GetAccountRankByPUUID(puuid, redis, client)
	if err != nil {
		fmt.Printf("Error getting valorant rank: %v\n", err)
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		response := &utils.DefaultResponse{
			Status: fasthttp.StatusInternalServerError,
			Data: &utils.MessageData{
				Message: err.Error(),
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}

	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	response := &utils.DefaultResponse{
		Status: fasthttp.StatusOK,
		Data:   rankData,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
