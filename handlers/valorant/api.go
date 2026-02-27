package valorant

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jckli/api/utils"
	"github.com/rueian/rueidis"
	"github.com/valyala/fasthttp"
)

const (
	hendrikBaseURL     = "https://api.henrikdev.xyz/valorant"
	officialContentURL = "https://valorant-api.com/v1/competitivetiers"
	officialAgentsURL  = "https://valorant-api.com/v1/agents"
	defaultRegion      = "na"
	defaultPlatform    = "pc"
)

func GetAccountRankByPUUID(puuid string, redis rueidis.Client, client *fasthttp.Client) (*HendrikMMRv3Data, error) {
	reqURL := fmt.Sprintf("%s/v3/by-puuid/mmr/%s/%s/%s", hendrikBaseURL, defaultRegion, defaultPlatform, puuid)
	resp, err := doValorantRequest(reqURL, redis, client)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode())
	}

	var wrapper HendrikMMRv3Response
	if err := json.Unmarshal(resp.Body(), &wrapper); err != nil {
		return nil, err
	}

	enrichTier(&wrapper.Data.Current.Tier, getTiersMap(redis, client))
	return &wrapper.Data, nil
}

func GetMatchesByPUUID(puuid string, redis rueidis.Client, client *fasthttp.Client) ([]EnrichedMatch, error) {
	reqURL := fmt.Sprintf("%s/v4/by-puuid/matches/%s/%s/%s", hendrikBaseURL, defaultRegion, defaultPlatform, puuid)
	resp, err := doValorantRequest(reqURL, redis, client)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(resp)

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode())
	}

	var rawWrapper HendrikMatchv4Response
	if err := json.Unmarshal(resp.Body(), &rawWrapper); err != nil {
		return nil, err
	}

	tiersMap := getTiersMap(redis, client)
	enrichedMatches := make([]EnrichedMatch, 0, len(rawWrapper.Data))

	for i := range rawWrapper.Data {
		for j := range rawWrapper.Data[i].Players {
			enrichTier(&rawWrapper.Data[i].Players[j].Tier, tiersMap)
		}
		enrichedMatches = append(enrichedMatches, EnrichedMatch{
			MatchV4Data: rawWrapper.Data[i],
			MyStats:     calculateMyStats(rawWrapper.Data[i], puuid, redis, client),
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

func getAgentIcon(agentID string, redis rueidis.Client, client *fasthttp.Client) (string, error) {
	var agentsResp OfficialAgentsResponse

	cached, err := utils.GetValorantAgentsCache(redis)
	if err == nil && cached != "" {
		if err := json.Unmarshal([]byte(cached), &agentsResp); err == nil {
			for _, agent := range agentsResp.Data {
				if strings.EqualFold(agent.UUID, agentID) {
					return agent.DisplayIcon, nil
				}
			}
		}
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(officialAgentsURL)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := client.Do(req, resp); err != nil {
		return "", err
	}

	if err := json.Unmarshal(resp.Body(), &agentsResp); err != nil {
		return "", err
	}

	if err := utils.SetValorantAgentsCache(redis, string(resp.Body())); err != nil {
		fmt.Printf("Warning: Failed to cache valorant agents: %v\n", err)
	}

	for _, agent := range agentsResp.Data {
		if strings.EqualFold(agent.UUID, agentID) {
			return agent.DisplayIcon, nil
		}
	}

	return "", fmt.Errorf("agent id %s not found", agentID)
}

func calculateMyStats(match MatchV4Data, myPUUID string, redis rueidis.Client, client *fasthttp.Client) DerivedStats {
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

	if totalRounds == 0 {
		totalRounds = 1
	}

	netDamage := me.Stats.Damage.Dealt - me.Stats.Damage.Received
	ddPerRound := float64(netDamage) / float64(totalRounds)
	acs := float64(me.Stats.Score) / float64(totalRounds)

	totalShots := me.Stats.Headshots + me.Stats.Bodyshots + me.Stats.Legshots
	hsPercent := 0.0
	if totalShots > 0 {
		hsPercent = (float64(me.Stats.Headshots) / float64(totalShots)) * 100
	}

	agentIcon, _ := getAgentIcon(me.Agent.ID, redis, client)

	return DerivedStats{
		Result:              result,
		Score:               fmt.Sprintf("%d-%d", roundsWon, roundsLost),
		Agent:               me.Agent.Name,
		AgentIconURL:        agentIcon,
		KDA:                 fmt.Sprintf("%d/%d/%d", me.Stats.Kills, me.Stats.Deaths, me.Stats.Assists),
		RankInGame:          me.Tier.Name,
		DamageDeltaPerRound: float64(int(ddPerRound*10)) / 10,
		ACS:                 float64(int(acs + 0.5)),
		HSPercent:           float64(int(hsPercent*10)) / 10,
	}
}

func getTiersMap(redis rueidis.Client, client *fasthttp.Client) map[int]TierVisuals {
	tierMap := make(map[int]TierVisuals)
	var tiersResp OfficialTiersResponse

	cached, err := utils.GetValorantTiersCache(redis)
	if err != nil || cached == "" || json.Unmarshal([]byte(cached), &tiersResp) != nil {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		req.SetRequestURI(officialContentURL)
		req.Header.SetMethod("GET")

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if client.Do(req, resp) == nil && json.Unmarshal(resp.Body(), &tiersResp) == nil {
			_ = utils.SetValorantTiersCache(redis, string(resp.Body()))
		}
	}

	if len(tiersResp.Data) > 0 {
		for _, t := range tiersResp.Data[len(tiersResp.Data)-1].Tiers {
			c, bg := "#ffffff", "#000000"
			if len(t.Color) >= 6 {
				c = "#" + t.Color[:6]
			}
			if len(t.BackgroundColor) >= 6 {
				bg = "#" + t.BackgroundColor[:6]
			}
			tierMap[t.Tier] = TierVisuals{Icon: t.LargeIcon, Color: c, BgColor: bg}
		}
	}
	return tierMap
}

func enrichTier(tier *PlayerTier, tiersMap map[int]TierVisuals) {
	if v, exists := tiersMap[tier.ID]; exists {
		tier.RankIconURL, tier.RankColor, tier.RankBackgroundColor = v.Icon, v.Color, v.BgColor
	} else {
		tier.RankColor, tier.RankBackgroundColor = "#ffffff", "#000000"
	}
}
