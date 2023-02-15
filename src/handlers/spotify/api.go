package spotify

import (
	"github.com/valyala/fasthttp"
	"os"
	"fmt"
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func getSpotifyToken() (*SpotifyRefreshResponse, error)  {
	urls := "https://accounts.spotify.com/api/token"
	data := url.Values{}
	data.Set("refresh_token", os.Getenv("SPOTIFY_REFRESH_TOKEN"))
	data.Set("grant_type", "refresh_token")
	body := []byte(data.Encode())
	resp, err := postRequest(urls, body)
	if err != nil {
		return nil, err
	}
	
	respBody := &SpotifyRefreshResponse{}
	if err = json.Unmarshal(resp, respBody); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return respBody, nil
}

func getTopItems(access_token, item_type, time_range string, limit int) (*SpotifyTopItemsResponse, error) {
	urls := "https://api.spotify.com/v1/me/top/" + item_type + "?time_range=" + time_range + "&limit=" + fmt.Sprintf("%v", limit)
	resp, err := getRequest(urls, access_token)
	if err != nil {
		return nil, err
	}

	respBody := &SpotifyTopItemsResponse{}
	if err = json.Unmarshal(resp, respBody); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return respBody, nil
}

func getRequest(url, access_token string) ([]byte, error) {
	client := &fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", "Bearer "+ access_token)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() == 401 {
		return nil, fmt.Errorf("Unauthorized")
	}

	return resp.Body(), nil
}

func postRequest(url string, body []byte) ([]byte, error) {
	client := &fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)

	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	auth := base64.StdEncoding.EncodeToString([]byte(client_id + ":" + client_secret))
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", "Basic "+auth)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Accept", "application/json")
    req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}
	if resp.StatusCode() == 401 {
		return nil, fmt.Errorf("Unauthorized")
	}

	return resp.Body(), nil
}