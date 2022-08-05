package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Links struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	ShortId      string             `json:"shortid"`
	Fullurl      string             `json:"fullurl"`
	Createdby    string             `json:"createdby"`
	Creationdate string             `json:"creationdate"`
	Counter      int                `json:"counter"` // would be a waste of resources to update every time - should implement in memory store that updates records every x hours etc
}
