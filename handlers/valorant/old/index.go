package valorant_old

import (
	"encoding/json"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := &DefaultResponse{
		Status: 200,
		Data: &MessageData{
			Message: "jckli api v2 - valorant endpoint",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
