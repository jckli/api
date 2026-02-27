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
	agentsMap := getAgentsMap(redis, client)
	enrichedMatches := make([]EnrichedMatch, 0, len(rawWrapper.Data))

	for i, match := range rawWrapper.Data {
		teamRounds := make(map[string]int)
		for _, t := range match.Teams {
			r := t.Rounds.Won + t.Rounds.Lost
			if r == 0 {
				r = 1
			}
			teamRounds[t.TeamID] = r
		}

		for j, p := range match.Players {
			enrichTier(&rawWrapper.Data[i].Players[j].Tier, tiersMap)
			rawWrapper.Data[i].Players[j].Agent.IconURL = agentsMap[strings.ToLower(p.Agent.ID)]

			tr := teamRounds[p.TeamID]
			if tr == 0 {
				tr = 1
			}

			stats := &rawWrapper.Data[i].Players[j].Stats
			netDamage := stats.Damage.Dealt - stats.Damage.Received
			totalShots := stats.Headshots + stats.Bodyshots + stats.Legshots

			hsPercent := 0.0
			if totalShots > 0 {
				hsPercent = (float64(stats.Headshots) / float64(totalShots)) * 100
			}

			stats.KDA = fmt.Sprintf("%d/%d/%d", stats.Kills, stats.Deaths, stats.Assists)
			stats.DamageDeltaPerRound = float64(int((float64(netDamage)/float64(tr))*10)) / 10
			stats.ACS = float64(int((float64(stats.Score) / float64(tr)) + 0.5))
			stats.HSPercent = float64(int(hsPercent*10)) / 10
		}

		enrichedMatches = append(enrichedMatches, EnrichedMatch{
			MatchV4Data: rawWrapper.Data[i],
			MyStats:     extractMyStats(rawWrapper.Data[i], puuid),
		})
	}

	return enrichedMatches, nil
}

func extractMyStats(match MatchV4Data, myPUUID string) DerivedStats {
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

	result, scoreStr := "Draw", "0-0"

	for _, t := range match.Teams {
		if t.TeamID == me.TeamID {
			scoreStr = fmt.Sprintf("%d-%d", t.Rounds.Won, t.Rounds.Lost)
			if t.Won {
				result = "Victory"
			} else if t.Rounds.Won < t.Rounds.Lost {
				result = "Defeat"
			}
			break
		}
	}

	return DerivedStats{
		Result:              result,
		Score:               scoreStr,
		Agent:               me.Agent.Name,
		AgentIconURL:        me.Agent.IconURL,
		KDA:                 me.Stats.KDA,
		RankInGame:          me.Tier.Name,
		DamageDeltaPerRound: me.Stats.DamageDeltaPerRound,
		ACS:                 me.Stats.ACS,
		HSPercent:           me.Stats.HSPercent,
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

func getAgentsMap(redis rueidis.Client, client *fasthttp.Client) map[string]string {
	agentMap := make(map[string]string)
	var agentsResp OfficialAgentsResponse

	cached, err := utils.GetValorantAgentsCache(redis)
	if err != nil || cached == "" || json.Unmarshal([]byte(cached), &agentsResp) != nil {
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		req.SetRequestURI(officialAgentsURL)
		req.Header.SetMethod("GET")
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		if client.Do(req, resp) == nil && json.Unmarshal(resp.Body(), &agentsResp) == nil {
			_ = utils.SetValorantAgentsCache(redis, string(resp.Body()))
		}
	}

	for _, a := range agentsResp.Data {
		agentMap[strings.ToLower(a.UUID)] = a.DisplayIcon
	}
	return agentMap
}

func enrichTier(tier *PlayerTier, tiersMap map[int]TierVisuals) {
	if v, exists := tiersMap[tier.ID]; exists {
		tier.RankIconURL, tier.RankColor, tier.RankBackgroundColor = v.Icon, v.Color, v.BgColor
	} else {
		tier.RankColor, tier.RankBackgroundColor = "#ffffff", "#000000"
	}
}
