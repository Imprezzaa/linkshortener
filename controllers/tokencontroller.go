package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Imprezzaa/linkshortener/auth"
	"github.com/Imprezzaa/linkshortener/models"
	"github.com/Imprezzaa/linkshortener/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var request TokenRequest
	var user models.User

	// check to see if the fields sent in the request bind to the TokenRequest struct
	// request holds the information from the current request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error: please check that your request is formed correctly",
			Data:    map[string]interface{}{"data": err.Error()},
		})
		return
	}

	// check to see if the email address is in the DB and if so bind it to the user var
	// that is of type models.User
	err = userCollection.FindOne(ctx, bson.M{"email": request.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "error, email not found",
			Data:    map[string]interface{}{"data": err.Error()},
		})
		return
	}

	// check the password sent in the user request against the password hash pulled from the DB
	credErr := user.CheckPassword(request.Password)
	if credErr != nil {
		c.JSON(http.StatusUnauthorized, responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "error: invalid credentials",
			Data:    map[string]interface{}{"data": err.Error()},
		})
		return
	}

	// generate a token for the user
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error: internal error, please try again later",
			Data:    map[string]interface{}{"data": err.Error()},
		})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
