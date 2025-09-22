package mal

import (
	"encoding/json"
	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client, client *fasthttp.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := &utils.DefaultResponse{
		Status: 200,
		Data: &utils.MessageData{
			Message: "jckli api v1 - myanimelist endpoint",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
