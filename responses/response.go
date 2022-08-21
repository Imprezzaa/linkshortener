package responses

// TODO: revise, restructure Reponse structs

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LinkResponse struct {
	Status   int                    `json:"status"`
	Message  string                 `json:"message"`
	ShortURL string                 `json:"shorturl"`
	Data     map[string]interface{} `json:"data"`
}
