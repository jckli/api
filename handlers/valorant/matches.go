package valorant

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func MatchesHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client, client *fasthttp.Client) {
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

	matchesData, err := GetMatchesByPUUID(puuid, redis, client)
	if err != nil {
		fmt.Printf("Error getting valorant matches: %v\n", err)
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
		Data:   matchesData, // Returns the full API response structure
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
