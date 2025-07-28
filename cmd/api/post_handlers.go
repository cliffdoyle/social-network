package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cliffdoyle/social-network/internal/models"
)

// createPostHandler handles POST /posts requests
func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	//Maps to the incoming json from the client
	var input models.PostCreateInput

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//retrieve the userID from the request context
	authenticatedUser := r.Context().Value("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	post, err := app.postService.Create(ctx, input, authenticatedUser)
	if err != nil {
		if errors.Is(err, models.ErrInvalidFieldInput) {
			// This is likely a 422 Unprocessable Entity
			app.failedValidationResponse(w, r, map[string]any{"validation error": models.ErrInvalidFieldInput})
			return
		}
		if errors.Is(err, sql.ErrNoRows) {
			//record not found in the database
			app.notFoundResponse(w, r)
			return

		}
		//otherwise we return default error
		app.serverErrorResponse(w, r, err)
		return
	}

	//send data to client through json marshal
	err = app.writeJSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//getPostHandler handles GET /posts/{id} requests
func (app *application) getPostHandler(w http.ResponseWriter,r *http.Request){
	//read the postID from the request URL
	postID:=app.rea
}
