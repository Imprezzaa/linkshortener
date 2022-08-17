package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

/*
TODO:
password field or seperate model for user creation and authentication
translate fields to MongoDB primitives

*/

type User struct {
	Username     string             `json:"username,omitempty" validate:"required"`
	Password     string             `json:"password" validate:"required"`
	Email        string             `json:"email" validate:"required"`
	Creationdate primitive.DateTime `json:"creationdate,omitempty"`
}

// bcrypt cost set to low low level for testing purposes
// TODO: set to a recommended level or above
// Hash password takes in a plain text password and returns a bcrypt hash of the password
func (u *User) HashPassword(password string) error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return err
	}
	u.Password = string(passwd)
	return nil
}

// CheckPassword compares the password sent by the user against the previously stored hash of the password
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
