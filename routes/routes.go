package routes

import (
	"github.com/Imprezzaa/linkshortener/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/:shortId", controllers.GetLink())

	users := r.Group("/u")

	users.POST("/create", controllers.CreateUser())

	users.GET("/:userId", controllers.GetUser())
	users.PUT("/:userId", controllers.EditUser())
	users.DELETE("/:userId", controllers.DeleteUser())

	users.GET("/users", controllers.GetAllUsers())

	links := r.Group("/l")

	links.POST("/create", controllers.CreateLink())
	links.GET("/:username/links", controllers.GetUserLinks())
}
