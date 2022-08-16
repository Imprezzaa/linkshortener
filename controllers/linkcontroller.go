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
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var linkCollection *mongo.Collection = db.GetCollection(db.DB, "links")

func CreateLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var links models.Link

		err := c.BindJSON(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: could not create link",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		err = validate.Struct(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: incorrect format",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		shortLink := MakeString(8)
		newLink := models.Link{
			FullURL:      links.FullURL,
			ShortID:      shortLink,
			CreatedBy:    links.CreatedBy,
			CreationDate: GetTime(),
		}

		result, err := linkCollection.InsertOne(ctx, newLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{
				Status:  http.StatusInternalServerError,
				Message: "unable to create link, please try again later.",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.LinkResponse{
			Status:   http.StatusCreated,
			Message:  "Success! Your link has been shortened!",
			ShortURL: "http://localhost:8080/" + shortLink,
			Data:     map[string]interface{}{"data": result},
		})
	}
}

func GetLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shortid := c.Param("shortid")

		var results map[string]interface{}
		err := linkCollection.FindOne(ctx, bson.M{"shortid": shortid}).Decode(&results)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.LinkResponse{
				Status:  http.StatusNotFound,
				Message: "error: url not found",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}

func GetLinkRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shortid := c.Param("shortid")
		opts := options.FindOne().SetProjection(bson.D{{Key: "fullurl", Value: 1}})

		var results map[string]string
		err := linkCollection.FindOne(ctx, bson.M{"shortid": shortid}, opts).Decode(&results)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{
				Status:  http.StatusNotFound,
				Message: "error: url not found",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// TODO: for debugging, remove later
		for _, result := range results {
			fmt.Println(result)
		}

		location := url.URL{Path: results["fullurl"]}
		varLocation := VerifyURL(location.Path)

		c.Redirect(http.StatusMovedPermanently, varLocation)

	}
}

func EditLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		// create a context with timeout for when the request is sent to the DB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// create a variable of type models.Link
		var links models.Link

		// shortid holds the value of the short_id field from the user request
		shortid := c.Param("shortid")

		// check to see if the body of the request will bind to the Link struct model
		// takes a pointer to the links var
		err := c.BindJSON(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: please verify that your request is formed correctly",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// pretty sure this checks the fields of the struct to validate that it contains
		// correct data.
		err = validate.Struct(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: please verify that your request is formed correctly",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		filter := bson.M{"shortid": shortid}
		update := bson.M{"$set": bson.M{"fullurl": links.FullURL}}
		result, err := linkCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{
				Status:  http.StatusInternalServerError,
				Message: "error: something went wrong, please try again.",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// create a new var of type models.Link so that we can return the updated
		// data to the user on a succesful update
		var updatedLink models.Link
		if result.MatchedCount == 1 {
			err := linkCollection.FindOne(ctx, filter).Decode(&updatedLink)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LinkResponse{
					Status:  http.StatusInternalServerError,
					Message: "error: something went wrong, please try again.",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}

		c.JSON(http.StatusOK, responses.LinkResponse{
			Status:  http.StatusOK,
			Message: "Success! Your link has been updated!",
			Data:    map[string]interface{}{"data": updatedLink},
		})
	}
}

func DeleteLink() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var links models.Link

		shortid := c.Param("shortid")

		err := c.BindJSON(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: please verify that your request is formed correctly",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		err = validate.Struct(&links)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.LinkResponse{
				Status:  http.StatusBadRequest,
				Message: "error: please verify that your request is formed correctly",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		filter := bson.M{"shortid": shortid}
		result, err := linkCollection.DeleteOne(ctx, filter)
		if err != nil || result.DeletedCount == 0 {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{
				Status:  http.StatusInternalServerError,
				Message: "error: either the document wasn't found or something went wrong, please try again",
			})
			return
		}

		// using DeleteOne means DeletedCount is either 0 or 1
		c.JSON(http.StatusOK, responses.LinkResponse{
			Status:  http.StatusOK,
			Message: "The link has sucessfully been deleted!",
			Data:    map[string]interface{}{"data": nil},
		})
	}
}

// GetUserLinks returns every link that the user has created
func GetUserLinks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var links []models.Link
		userName := c.Param("userid")
		filter := bson.D{{Key: "createdby", Value: userName}}
		projection := bson.M{"fullurl": 1, "shortid": 1, "_id": 0}

		results, err := linkCollection.Find(ctx, filter, options.Find().SetProjection(projection))
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LinkResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// need to close the cursor at the end of the function
		// returns nearly correct data, just also sends a blank "createdby" field
		// TODO: bug fix
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleLink models.Link
			if err = results.Decode(&singleLink); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LinkResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
			}
			links = append(links, singleLink)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    map[string]interface{}{"data": links}},
		)
	}
}
