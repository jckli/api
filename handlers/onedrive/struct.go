package onedrive

type OnedriveTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OnedriveItemsResponse struct {
	Value []struct {
		Id                   string `json:"id"`
		Name                 string `json:"name"`
		WebUrl               string `json:"webUrl"`
		ParentId             string `json:"parentId"`
		LastModifiedDateTime string `json:"lastModifiedDateTime"`
		Size                 int    `json:"size"`
		DownloadUrl          string `json:"@microsoft.graph.downloadUrl"`
		File                 struct {
			MimeType string `json:"mimeType"`
		} `json:"file"`
		Folder struct {
			ChildCount int `json:"childCount"`
		} `json:"folder"`
	} `json:"value"`
}
