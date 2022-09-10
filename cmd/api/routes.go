package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/v1/links", app.createLinkHandler)
	router.HandlerFunc(http.MethodGet, "/v1/links/:shortid", app.showLinkHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/links/:shortid", app.patchLinkHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/links/:shortid", app.deleteLinkHandler)

	return router
}
