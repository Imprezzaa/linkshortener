package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// reorganized shortLink struct
// Removed counter, will later reimplement in a different struct or package dealing specifically with counting the links
/*
type ShortLink struct {
	ShortID      string `json:"short_id,omitempty"`                     // not required, it is provided by the controller package when the link is submitted
	LongURL      string `json:"long_url,omitempty" validate:"required"` // required by user
	CreatedBy    string `json:"created_by" validate:"required"`         // should be sent with the request to create a new link, should contain a username
	Creationdate int64  `json:"creation_date,omitempty"`                // not required by the user, is added by the controller
	Counter      uint32 `json:"counter,omitempty"`                      // would be a waste of resources to update every time - should implement in memory store that updates records every x hours etc
}
*/

// Notes:
// x64 wordsize is 8
// strings are 2*wordsize = 16 bytes
// int64 is 8 bytes/1 word
// uint32 is 4 bytes, half a word. if last field in struct it adds 4 bytes of padding?
// not so sure about

// NewLink outlines the structure that user requests should follow
// CreateedBy and CreationDate are handled by the link_controller
type Link struct {
	FullURL      string             `json:"full_url" validate:"required"`
	ShortID      string             `json:"short_id,omitempty"`
	CreatedBy    string             `json:"created_by" validate:"required"`
	CreationDate primitive.DateTime `json:"creation_date,omitempty"`
}
