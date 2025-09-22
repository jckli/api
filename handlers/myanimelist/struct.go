package mal

type MalTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MalMangaListResponse struct {
	Data   []MalMangaEntry `json:"data"`
	Paging Paging          `json:"paging"`
}

type MalMangaEntry struct {
	Node       MalMangaNode       `json:"node"`
	ListStatus MalMangaListStatus `json:"list_status"`
}

type MalMangaNode struct {
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
