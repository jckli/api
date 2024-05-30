package onedrive

type OnedriveTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OnedriveItemsResponse struct {
	Value []struct {
		Name                 string `json:"name"`
		LastModifiedDateTime string `json:"lastModifiedDateTime"`
		Size                 int    `json:"size"`
		DownloadUrl          string `json:"@microsoft.graph.downloadUrl"`
		File                 struct {
			MimeType string `json:"mimeType"`
		} `json:"file"`
		Image struct {
			Height int `json:"height"`
			Width  int `json:"width"`
		} `json:"image"`
	} `json:"value"`
}
