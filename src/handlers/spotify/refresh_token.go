package spotify

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/rueian/rueidis"
)

func RefreshTokenHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	access_token, err := getSpotifyToken()
	if err != nil {
		ctx.Response.SetStatusCode(401)
		response := &DefaultResponse{
			Status: 401,
			Data: &MessageData{
				Message: "No access.",
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
		Data: access_token,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}