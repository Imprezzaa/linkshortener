package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Notes:
// x64 wordsize is 8
// strings are 2*wordsize = 16 bytes
// int64 is 8 bytes/1 word
// uint32 is 4 bytes, half a word. if last field in struct it adds 4 bytes of padding?
// not so sure about

// Link outlines the structure that user requests should follow
// CreateedBy and CreationDate are handled by the link_controller
type Link struct {
	FullURL      string             `json:"fullurl" validate:"required"`
	ShortID      string             `json:"shortid,omitempty"`
	Username     string             `json:"username" validate:"required"`
	CreationDate primitive.DateTime `json:"creationdate,omitempty"`
}
