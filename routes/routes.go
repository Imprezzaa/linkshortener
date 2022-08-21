package routes

import (
	"github.com/Imprezzaa/linkshortener/controllers"
	"github.com/gin-gonic/gin"
)

// TODO: Revise, regroup routes and add some protected routes using Auth middleware

func Routes(r *gin.Engine) {

	r.GET("/:shortid", controllers.GetLinkRedirect())

	users := r.Group("/u")

	users.POST("/create", controllers.CreateUser())

	users.GET("/:userid", controllers.GetUser())
	users.PUT("/:userid", controllers.EditUser())
	users.DELETE("/:userid", controllers.DeleteUser())

	users.GET("/users", controllers.GetAllUsers())
	users.GET("/:userid/links", controllers.GetUserLinks())

	links := r.Group("/l")

	links.POST("/create", controllers.CreateLink())

	links.GET("/:shortid", controllers.GetLink())
	links.PUT("/:shortid", controllers.EditLink())
	links.DELETE("/:shortid", controllers.DeleteLink())

}
