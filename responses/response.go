package responses

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LinkResponse struct {
	Status   int                    `json:"status"`
	Message  string                 `json:"message"`
	ShortURL string                 `json:"short_url"`
	Data     map[string]interface{} `json:"data"`
}
