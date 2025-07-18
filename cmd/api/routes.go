package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/update-privacy", app.updatePrivacyHandler)
	// mux.HandleFunc("/healthcheck", app.healthcheckHandler,app.registerUserHandler)
	// mux.HandleFunc("/test", app.errorTest)
	// mux.HandleFunc("/update-privacy", )

	// mux.HandleFunc("/register", ) // Simple /register endpoint

	return app.rateLimit(router)
}
