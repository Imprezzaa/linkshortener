package main

import (
	"github.com/Imprezzaa/linkshortener/db"
	"github.com/Imprezzaa/linkshortener/routes"
	"github.com/gin-gonic/gin"
)

/*
TODO:
https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
make service more resilient in the case of a shutdown
*/

func main() {
	r := gin.Default()

	// connects to the database
	db.ConnectDB()

	// Loads the routes for user and link logic
	routes.Routes(r)

	r.Run("localhost:8080")
}
