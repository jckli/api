package valorant

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/rueian/rueidis"
)

func IndexHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := &DefaultResponse{
		Status: 200,
		Data: &MessageData{
			Message: "jckli api v1 - valorant endpoint",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}