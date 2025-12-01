package valorant

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

const (
	hendrikBaseURL  = "https://api.henrikdev.xyz/valorant"
	defaultRegion   = "na"
	defaultPlatform = "pc"
)

func GetAccountRankByPUUID(puuid string, redis rueidis.Client, client *fasthttp.Client) (*HendrikMMRv3Data, error) {
	reqURL := fmt.Sprintf(
		"%s/v3/by-puuid/mmr/%s/%s/%s",
		hendrikBaseURL,
		defaultRegion,
		defaultPlatform,
		puuid,
	)

	resp, err := doValorantRequest(reqURL, redis, client)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	wrapper := &HendrikMMRv3Response{}
	if err := json.Unmarshal(resp.Body(), &wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rank response: %w", err)
	}

	return &wrapper.Data, nil
}

func GetMatchesByPUUID(puuid string, redis rueidis.Client, client *fasthttp.Client) ([]MatchV4Data, error) {
	reqURL := fmt.Sprintf(
		"%s/v4/by-puuid/matches/%s/%s/%s",
		hendrikBaseURL,
		defaultRegion,
		defaultPlatform,
		puuid,
	)

	resp, err := doValorantRequest(reqURL, redis, client)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d, body: %s", resp.StatusCode(), string(resp.Body()))
	}

	wrapper := &HendrikMatchv4Response{}
	if err := json.Unmarshal(resp.Body(), &wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal matches response: %w", err)
	}

	return wrapper.Data, nil
}

func doValorantRequest(url string, redis rueidis.Client, client *fasthttp.Client) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	if apiKey := os.Getenv("VALORANT_API_KEY"); apiKey != "" {
		req.Header.Set("Authorization", apiKey)
	}

	resp := fasthttp.AcquireResponse()
	if err := client.Do(req, resp); err != nil {
		fasthttp.ReleaseResponse(resp)
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}
