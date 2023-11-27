package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/play_tents", app.requirePermission("play_tents:read", app.listPlayTentsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/play_tents", app.requirePermission("play_tents:write", app.createPlayTentHandler))
	router.HandlerFunc(http.MethodGet, "/v1/play_tents/:id", app.requirePermission("play_tents:read", app.showPlayTentHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/play_tents/:id", app.requirePermission("play_tents:write", app.updatePlayTentHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/play_tents/:id", app.requirePermission("play_tents:write", app.deletePlayTentHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
