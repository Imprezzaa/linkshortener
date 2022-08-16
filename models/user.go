package models

/*
TODO:
password field or seperate model for user creation and authentication
translate fields to MongoDB primitives

*/

type User struct {
	UserName     string `json:"username,omitempty" validate:"required"`
	Creationdate int64  `json:"creationdate,omitempty"`
}
