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
	router.HandlerFunc(http.MethodPost, "/login", app.LoginHandler)

	router.Handler(http.MethodPost, "/v1/create-post", app.Authenticator(http.HandlerFunc(app.createPostHandler)))
	router.Handler(http.MethodGet, "/get-posts/{id}", app.Authenticator(http.HandlerFunc(app.getPostHandler)))
	router.Handler(http.MethodPatch, "/update-post/{id}", app.Authenticator(http.HandlerFunc(app.updatePostHandler)))
	router.Handler(http.MethodDelete, "/delete-post/{id}", app.Authenticator(http.HandlerFunc(app.deletePostHandler)))
	return app.rateLimit(router)
}
