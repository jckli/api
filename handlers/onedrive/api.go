package onedrive

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
	"os"
)

func getOnedriveToken(
	redis rueidis.Client,
	client *fasthttp.Client,
) (string, error) {
	tokens, err := utils.GetOnedriveRedisTokens(redis)
	if err != nil {
		return "", err
	}

	refreshToken := tokens.RefreshToken
	if tokens.AccessToken != "" {
		return tokens.AccessToken, nil
	}

	if tokens.RefreshToken == "" {
		refreshToken = os.Getenv("ONEDRIVE_REFRESH_TOKEN")
	}

	url := "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	args := &fasthttp.Args{}
	args.Add("client_id", os.Getenv("ONEDRIVE_CLIENT_ID"))
	args.Add("client_secret", os.Getenv("ONEDRIVE_CLIENT_SECRET"))
	args.Add("grant_type", "refresh_token")
	args.Add("refresh_token", refreshToken)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	req.SetBody(args.QueryString())

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return "", err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf(
			"Non-200 response status: %d, body: %s",
			resp.StatusCode(),
			resp.Body(),
		)
	}

	respData := &OnedriveTokenResponse{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		return "", fmt.Errorf("Failed to parse JSON response: %v", err)
	}

	newTokens := &utils.OnedriveTokens{
		RefreshToken:      respData.RefreshToken,
		AccessToken:       respData.AccessToken,
		AccessTokenExpiry: respData.ExpiresIn,
	}

	if err := utils.SetOnedriveRedisTokens(redis, newTokens); err != nil {
		return "", err
	}

	return respData.AccessToken, nil
}

func getFolderItems(
	redis rueidis.Client,
	client *fasthttp.Client,
	folderId string,
) (*OnedriveItemsResponse, error) {
	access_token, err := getOnedriveToken(redis, client)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"https://graph.microsoft.com/v1.0/me/drive/items/%s/children?$expand=thumbnails",
		folderId,
	)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+access_token)
	req.URI().
		QueryArgs().
		Set("select", "@microsoft.graph.downloadUrl,name,file,size,lastModifiedDateTime,image")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf(
			"Non-200 response status: %d, body: %s",
			resp.StatusCode(),
			resp.Body(),
		)
	}

	respData := &OnedriveItemsResponse{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON response: %v", err)
	}

	return respData, nil
}
