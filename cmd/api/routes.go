package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/play_tents", app.listPlayTentsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/play_tents", app.createPlayTentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/play_tents/:id", app.showPlayTentHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/play_tents/:id", app.updatePlayTentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/play_tents/:id", app.deletePlayTentHandler)

	return router
}
