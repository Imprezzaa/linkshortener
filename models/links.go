package models

type ShortLink struct {
	ShortID      string `json:"shortid,omitempty"`
	LongURL      string `json:"long_url,omitempty" validate:"required"`
	CreatedBy    string `json:"createdby" validate:"required"`
	Creationdate int64  `json:"creationdate,omitempty"`
	Counter      int    `json:"counter,omitempty"` // would be a waste of resources to update every time - should implement in memory store that updates records every x hours etc
}
