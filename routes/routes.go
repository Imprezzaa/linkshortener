package routes

import (
	"github.com/Imprezzaa/linkshortener/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/:shortId", controllers.GetLink())

	users := r.Group("/u")

	users.POST("/create", controllers.CreateUser())

	users.GET("/:username", controllers.GetUser())
	users.PUT("/:username", controllers.EditUser())
	users.DELETE("/:username", controllers.DeleteUser())

	users.GET("/users", controllers.GetAllUsers())

	links := r.Group("/l")

	// create a short link
	links.POST("/create", controllers.CreateLink())

	// Find single link by its short_id
	links.GET("/:short_id", controllers.GetLink())

	// update a short link by its short_id
	links.PUT("/:short_id")

	// delete a link based on it's short_id
	links.DELETE("/:short_id")

	links.GET("/:username/links", controllers.GetUserLinks()) // should probably move this to the user group??
}
