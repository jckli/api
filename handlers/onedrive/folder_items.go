package onedrive

import (
	"encoding/json"
	"fmt"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"

	"github.com/jckli/api/utils"
)

func FolderItemsHandler(ctx *fasthttp.RequestCtx, redis rueidis.Client, client *fasthttp.Client) {
	ctx.Response.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	folderId := ctx.UserValue("folderId").(string)

	// scuffed security system lmao only one folder needs to be reached for my website
	if folderId != "01NV5GJPTEVQP3Z732KJF2EWCKNROO7U7R" {
		folderId = "gg"
	}

	folderItems, err := getFolderItems(redis, client, folderId)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		response := &utils.DefaultResponse{
			Status: 500,
			Data: &utils.MessageData{
				Message: err.Error(),
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}

	if folderItems == nil {
		ctx.Response.SetStatusCode(404)
		response := &utils.DefaultResponse{
			Status: 404,
			Data: &utils.MessageData{
				Message: fmt.Sprintf("Folder (%s) not found.", folderId),
			},
		}
		if err := json.NewEncoder(ctx).Encode(response); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
		return
	}

	ctx.Response.SetStatusCode(200)
	response := &utils.DefaultResponse{
		Status: 200,
		Data:   folderItems,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
