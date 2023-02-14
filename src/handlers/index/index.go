package index

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	ctx.Response.SetStatusCode(200)
	response := &IndexResponse{
		Status: 200,
		Data: &IndexData{
			Message: "jckli api v1",
			Link: "https://jackli.dev",
		},
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}