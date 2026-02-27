package valorant

type PlayerTier struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	RankIconURL         string `json:"rank_icon_url,omitempty"`
	RankColor           string `json:"rank_color,omitempty"`
	RankBackgroundColor string `json:"rank_background_color,omitempty"`
}

type HendrikMMRv3Response struct {
	Status int              `json:"status"`
	Data   HendrikMMRv3Data `json:"data"`
}

type HendrikMMRv3Data struct {
	Account struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	} `json:"account"`
	Current struct {
		Tier PlayerTier `json:"tier"`
		RR   int        `json:"rr"`
		Elo  int        `json:"elo"`
	} `json:"current"`
}

type HendrikMatchv4Response struct {
	Status int           `json:"status"`
	Data   []MatchV4Data `json:"data"`
}

type MatchV4Data struct {
	Metadata MatchV4Metadata `json:"metadata"`
	Players  []MatchV4Player `json:"players"`
	Teams    []MatchV4Team   `json:"teams"`
}

type MatchV4Metadata struct {
	Map struct {
		Name string `json:"name"`
	} `json:"map"`
	StartedAt string `json:"started_at"`
	Queue     struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		ModeType string `json:"mode_type"`
	} `json:"queue"`
	Cluster string `json:"cluster"`
}

type MatchV4Player struct {
	PUUID   string `json:"puuid"`
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	TeamID  string `json:"team_id"`
	PartyID string `json:"party_id"`
	Agent   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"agent"`
	Stats MatchV4Stats `json:"stats"`
	Tier  PlayerTier   `json:"tier"`
}

type MatchV4Stats struct {
	Score     int `json:"score"`
	Kills     int `json:"kills"`
	Deaths    int `json:"deaths"`
	Assists   int `json:"assists"`
	Headshots int `json:"headshots"`
	Bodyshots int `json:"bodyshots"`
	Legshots  int `json:"legshots"`
	Damage    struct {
		Dealt    int `json:"dealt"`
		Received int `json:"received"`
	} `json:"damage"`
}

type MatchV4Team struct {
	TeamID string `json:"team_id"`
	Won    bool   `json:"won"`
	Rounds struct {
		Won  int `json:"won"`
		Lost int `json:"lost"`
	} `json:"rounds"`
}

type EnrichedMatch struct {
	MatchV4Data
	MyStats DerivedStats `json:"my_stats"`
}

type DerivedStats struct {
	Result              string  `json:"result"`
	Score               string  `json:"score"`
	Agent               string  `json:"agent"`
	AgentIconURL        string  `json:"agent_icon_url"`
	KDA                 string  `json:"kda"`
	RankInGame          string  `json:"rank_in_game"`
	DamageDeltaPerRound float64 `json:"damage_delta_per_round"`
	ACS                 float64 `json:"acs"`
	HSPercent           float64 `json:"hs_percent"`
}

type OfficialTiersResponse struct {
	Data []struct {
		Tiers []struct {
			Tier            int    `json:"tier"`
			LargeIcon       string `json:"largeIcon"`
			Color           string `json:"color"`
			BackgroundColor string `json:"backgroundColor"`
		} `json:"tiers"`
	} `json:"data"`
}

type OfficialAgentsResponse struct {
	Data []struct {
		UUID        string `json:"uuid"`
		DisplayName string `json:"displayName"`
		DisplayIcon string `json:"displayIcon"`
	} `json:"data"`
}

type TierVisuals struct {
	Icon, Color, BgColor string
}
