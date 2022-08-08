package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Imprezzaa/linkshortener/db"
	"github.com/Imprezzaa/linkshortener/models"
	"github.com/Imprezzaa/linkshortener/responses"
	"github.com/Imprezzaa/linkshortener/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var LinkCollection *mongo.Collection = db.GetCollection(db.DB, "links")

func CreateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var links models.ShortLink

		if err := c.BindJSON(&links); err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{Status: http.StatusBadRequest, Message: "error: could not create link", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&links); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{Status: http.StatusBadRequest, Message: "error: incorrect format", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		shortLink := utils.MakeString(8)
		newLink := models.ShortLink{
			ShortID:      shortLink,
			LongURL:      links.LongURL,
			CreatedBy:    links.CreatedBy,
			Creationdate: time.Now().Unix(),
			Counter:      0,
		}

		result, err := LinkCollection.InsertOne(ctx, newLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "unable to create link, please try again later.", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.LinkResponse{Status: http.StatusCreated, Message: "Success! Your link has been shortened!", ShortURL: "http://localhost:8080/l/" + shortLink, Data: map[string]interface{}{"data": result}})
	}
}

func GetLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shorty := c.Param("shortId")
		opts := options.FindOne().SetProjection(bson.D{{Key: "longurl", Value: 1}})

		var results map[string]string
		err := LinkCollection.FindOne(ctx, bson.M{"shortid": shorty}, opts).Decode(&results)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error: url not found", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// TODO: for debugging, remove later
		for _, result := range results {
			fmt.Println(result)
		}

		location := url.URL{Path: results["longurl"]}
		varLocation := utils.VerifyURL(location.Path)

		c.Redirect(http.StatusMovedPermanently, varLocation)

	}
}

// GetUserLinks returns every link that the user has created
func GetUserLinks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var links []models.ShortLink
		userName := c.Param("username")
		filter := bson.D{{Key: "createdby", Value: userName}}
		projection := bson.M{"longurl": 1, "shortid": 1, "_id": 0}

		results, err := LinkCollection.Find(ctx, filter, options.Find().SetProjection(projection))
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// need to close the cursor at the end of the function
		// returns nearly correct data, just also sends a blank "createdby" field
		// TODO: bug fix
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleLink models.ShortLink
			if err = results.Decode(&singleLink); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LinkResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			links = append(links, singleLink)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": links}},
		)
	}
}
