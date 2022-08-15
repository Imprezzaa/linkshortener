package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
TODO:
password field or seperate model for user creation and authentication
translate fields to MongoDB primitives

should mongo db _id be used as a UUID or should a new UUID be generated for each user
*/

type User struct {
	UserName     string             `json:"username" validate:"required"`
	CreationDate primitive.DateTime `json:"creationDate,omitempty"`
	Location     string             `json:"location"`
}
