package mal

import (
	"encoding/json"
	"fmt"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"

	"github.com/jckli/api/utils"
)

func AnimeListHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client, client *fasthttp.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	mangaList, err := GetUserAnimeList(redis, client)
	if err != nil {
		fmt.Printf("Error getting user anime list: %v\n", err)
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
		Data:   mangaList,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
