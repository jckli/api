package mal

import "time"

type MalTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MalUnifiedListResponse struct {
	Data []MalUnifiedListEntry `json:"data"`
}

type MalUnifiedListEntry struct {
	Type       string         `json:"type"`
	UpdatedAt  time.Time      `json:"updated_at"`
	AnimeEntry *MalAnimeEntry `json:"anime_entry,omitempty"`
	MangaEntry *MalMangaEntry `json:"manga_entry,omitempty"`
}

type MalMangaListResponse struct {
	Data   []MalMangaEntry `json:"data"`
	Paging Paging          `json:"paging"`
}

type MalMangaEntry struct {
	Node       MalNode            `json:"node"`
	ListStatus MalMangaListStatus `json:"list_status"`
}

type MalAnimeListResponse struct {
	Data   []MalAnimeEntry `json:"data"`
	Paging Paging          `json:"paging"`
}

type MalAnimeEntry struct {
	Node       MalNode            `json:"node"`
	ListStatus MalAnimeListStatus `json:"list_status"`
}

type MalNode struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	MainPicture MalMainPicture `json:"main_picture"`
}

type MalMainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type MalMangaListStatus struct {
	Status          string `json:"status"`
	Score           int    `json:"score"`
	NumChaptersRead int    `json:"num_chapters_read"`
	NumVolumesRead  int    `json:"num_volumes_read"`
	IsRereading     bool   `json:"is_rereading"`
	UpdatedAt       string `json:"updated_at"`
}

type MalAnimeListStatus struct {
	Status             string `json:"status"`
	Score              int    `json:"score"`
	NumEpisodesWatched int    `json:"num_episodes_watched"`
	IsRewatching       bool   `json:"is_rewatching"`
	UpdatedAt          string `json:"updated_at"`
}

type MalAuthorEntry struct {
	Node MalAuthorNode `json:"node"`
	Role string        `json:"role"`
}

type MalAuthorNode struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Paging struct {
	Next string `json:"next"`
}
