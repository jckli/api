package spotify

type DefaultResponse struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type SpotifyRefreshResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
}

type SpotifyTopItemsResponse struct {
	Href string `json:"href"`
	Limit int `json:"limit"`
	Next string `json:"next"`
	Offset int `json:"offset"`
	Previous string `json:"previous"`
	Total int `json:"total"`
	Items []SpotifyTopItem `json:"items"`
}

type SpotifyTopItem struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href string `json:"href"`
		Total int `json:"total"`
	} `json:"followers"`
	Genres []string `json:"genres"`
	Href string `json:"href"`
	Id string `json:"id"`
	Images []struct {
		Url string `json:"url"`
		Height int `json:"height"`
		Width int `json:"width"`
	} `json:"images"`
	Name string `json:"name"`
	Popularity int `json:"popularity"`
	Type string `json:"type"`
	Uri string `json:"uri"`
}