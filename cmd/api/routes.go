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
	router.Handler(http.MethodGet, "/get-posts/:id", app.Authenticator(http.HandlerFunc(app.getPostHandler)))
	router.Handler(http.MethodPost, "/follow/:id", app.Authenticator(http.HandlerFunc(app.Follow)))
	router.Handler(http.MethodGet, "/followers", app.Authenticator(http.HandlerFunc(app.Followers)))
	router.Handler(http.MethodGet, "/following", app.Authenticator(http.HandlerFunc(app.Following)))
	router.Handler(http.MethodPost, "/unfollow/:id", app.Authenticator(http.HandlerFunc(app.Unfollow)))
	router.Handler(http.MethodGet, "/posts", app.Authenticator(http.HandlerFunc(app.PostsFeedByPrivacy)))
	router.Handler(http.MethodPatch, "/update-post/:id", app.Authenticator(http.HandlerFunc(app.updatePostHandler)))
	router.Handler(http.MethodDelete, "/delete-post/:id", app.Authenticator(http.HandlerFunc(app.deletePostHandler)))
	// Reaction endpoints
	router.Handler(http.MethodPost, "/api/posts/reactions", app.Authenticator(http.HandlerFunc(app.reactionHandler.ReactToPost)))
	router.Handler(http.MethodPost, "/api/comments/reactions", app.Authenticator(http.HandlerFunc(app.reactionHandler.ReactToComment)))
	return app.rateLimit(router)
}
