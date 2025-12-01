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

func GetMatchesByPUUID(puuid string, redis rueidis.Client, client *fasthttp.Client) ([]EnrichedMatch, error) {
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

	rawWrapper := &HendrikMatchv4Response{}
	if err := json.Unmarshal(resp.Body(), &rawWrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal matches response: %w", err)
	}

	enrichedMatches := make([]EnrichedMatch, 0, len(rawWrapper.Data))

	for _, match := range rawWrapper.Data {
		stats := calculateMyStats(match, puuid)

		enrichedMatches = append(enrichedMatches, EnrichedMatch{
			MatchV4Data: match,
			MyStats:     stats,
		})
	}

	return enrichedMatches, nil
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

func calculateMyStats(match MatchV4Data, myPUUID string) DerivedStats {
	var me MatchV4Player
	found := false

	for _, p := range match.Players {
		if p.PUUID == myPUUID {
			me = p
			found = true
			break
		}
	}

	if !found {
		return DerivedStats{Result: "Unknown"}
	}

	result := "Draw"
	roundsWon, roundsLost := 0, 0
	totalRounds := 1

	for _, t := range match.Teams {
		if t.TeamID == me.TeamID {
			roundsWon = t.Rounds.Won
			roundsLost = t.Rounds.Lost
			totalRounds = roundsWon + roundsLost

			if t.Won {
				result = "Victory"
			} else if roundsWon < roundsLost {
				result = "Defeat"
			}
			break
		}
	}

	netDamage := me.Stats.Damage.Dealt - me.Stats.Damage.Received
	ddPerRound := float64(netDamage) / float64(totalRounds)

	acs := 0.0
	if totalRounds > 0 {
		acs = float64(me.Stats.Score) / float64(totalRounds)
	}

	totalShots := me.Stats.Headshots + me.Stats.Bodyshots + me.Stats.Legshots
	hsPercent := 0.0
	if totalShots > 0 {
		hsPercent = (float64(me.Stats.Headshots) / float64(totalShots)) * 100
	}

	return DerivedStats{
		Result:              result,
		Score:               fmt.Sprintf("%d-%d", roundsWon, roundsLost),
		Agent:               me.Agent.Name,
		KDA:                 fmt.Sprintf("%d/%d/%d", me.Stats.Kills, me.Stats.Deaths, me.Stats.Assists),
		RankInGame:          me.Tier.Name,
		DamageDeltaPerRound: float64(int(ddPerRound*10)) / 10,
		ACS:                 float64(int(acs*10)) / 10,
		HSPercent:           float64(int(hsPercent*10)) / 10,
	}
}
