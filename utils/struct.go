package utils

type DefaultResponse struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type OnedriveTokens struct {
	AccessToken       string
	RefreshToken      string
	AccessTokenExpiry int64
}

type MalTokens struct {
	AccessToken  string
	RefreshToken string
}
