package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// user endpoints
	// create

	// get

	// patch

	// delete

	return router
}
