package valorant

type DefaultResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}
