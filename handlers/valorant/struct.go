package valorant

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
		Tier struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"tier"`
		RR  int `json:"rr"`
		Elo int `json:"elo"`
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
	PUUID  string `json:"puuid"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	TeamID string `json:"team_id"`
	Agent  struct {
		Name string `json:"name"`
	} `json:"agent"`
	Stats MatchV4Stats `json:"stats"`
	Tier  struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tier"`
}

type MatchV4Stats struct {
	Score   int `json:"score"`
	Kills   int `json:"kills"`
	Deaths  int `json:"deaths"`
	Assists int `json:"assists"`
}

type MatchV4Team struct {
	TeamID string `json:"team_id"`
	Won    bool   `json:"won"`
	Rounds struct {
		Won  int `json:"won"`
		Lost int `json:"lost"`
	} `json:"rounds"`
}
