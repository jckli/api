package utils

type DefaultResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type OnedriveTokens struct {
	AccessToken       string
	RefreshToken      string
	AccessTokenExpiry int64
}
