package mal

import (
	"encoding/json"
	"fmt"
	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

const (
	apiBaseURL = "https://api.myanimelist.net/v2"
	tokenURL   = "https://myanimelist.net/v1/oauth2/token"
)

func doMalRequest(req *fasthttp.Request, redis rueidis.Client, client *fasthttp.Client) (*fasthttp.Response, error) {
	accessToken, err := getMalToken(redis, client)
	if err != nil {
		return nil, fmt.Errorf("failed to get initial token: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		fasthttp.ReleaseResponse(resp)
		return nil, fmt.Errorf("failed to execute initial request: %w", err)
	}
	log.Printf("Initial Request Raw Response: %s", string(resp.Body()))

	if resp.StatusCode() == 401 {
		fasthttp.ReleaseResponse(resp)

		tokens, err := utils.GetMalRedisTokens(redis)
		if err != nil {
			return nil, fmt.Errorf("could not get tokens from redis for retry: %w", err)
		}

		newAccessToken, err := refreshAccessToken(redis, client, tokens.RefreshToken)
		if err != nil {
			return nil, fmt.Errorf("failed to refresh token: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+newAccessToken)
		resp = fasthttp.AcquireResponse()
		if err := client.Do(req, resp); err != nil {
			fasthttp.ReleaseResponse(resp)
			return nil, fmt.Errorf("failed to execute retry request: %w", err)
		}
	}

	return resp, nil
}

func GetUserMangaList(redis rueidis.Client, client *fasthttp.Client) (*MalMangaListResponse, error) {
	reqURL := fmt.Sprintf(
		"%s/users/@me/mangalist?sort=list_updated_at&limit=10&fields=list_status",
		apiBaseURL,
	)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(reqURL)
	req.Header.SetMethod("GET")

	resp, err := doMalRequest(req, redis, client)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code from MAL API: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	respData := &MalMangaListResponse{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal manga list response: %w", err)
	}

	return respData, nil
}

func getMalToken(redis rueidis.Client, client *fasthttp.Client) (string, error) {
	tokens, err := utils.GetMalRedisTokens(redis)
	if err != nil {
		return "", fmt.Errorf("error getting tokens from redis: %w", err)
	}

	if tokens.AccessToken == "" {
		return refreshAccessToken(redis, client, tokens.RefreshToken)
	}

	return tokens.AccessToken, nil
}

func refreshAccessToken(redis rueidis.Client, client *fasthttp.Client, refreshToken string) (string, error) {
	if refreshToken == "" {
		if refreshToken = os.Getenv("MAL_REFRESH_TOKEN"); refreshToken == "" {
			return "", fmt.Errorf("cannot refresh: MAL refresh token is missing")
		}
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(tokenURL)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.PostArgs().Add("client_id", os.Getenv("MAL_CLIENT_ID"))
	req.PostArgs().Add("client_secret", os.Getenv("MAL_CLIENT_SECRET"))
	req.PostArgs().Add("grant_type", "refresh_token")
	req.PostArgs().Add("refresh_token", refreshToken)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return "", fmt.Errorf("failed to execute token refresh request: %w", err)
	}

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("token refresh failed with status code %d: %s", resp.StatusCode(), string(resp.Body()))
	}

	respData := &MalTokenResponse{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		return "", fmt.Errorf("Failed to parse JSON response: %v", err)
	}

	newTokens := &utils.MalTokens{
		RefreshToken: respData.RefreshToken,
		AccessToken:  respData.AccessToken,
	}

	if err := utils.SetMalRedisTokens(redis, newTokens); err != nil {
		return "", err
	}

	fmt.Println("Successfully refreshed MAL tokens.")
	return respData.AccessToken, nil
}
