package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/validator"
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

// getPostHandler handles GET /posts/{id} requests
func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	//read the postID from the request URL
	postID := app.readIDParam(r)

	//context to determine maximum time requests take to fetch from db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	post, err := app.postService.GetByID(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	//send data to client through json marshal
	err = app.writeJSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//updatePostHandler handles PATCH /posts/{id} requests

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve post id from request url
	postID := app.readIDParam(r)

	//DTO to old post json from client temporarily
	var input models.PostUpdateInput

	//read incoming json to input struct
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//request context with timeout to determine time request takes
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//call the service layer to handle validation of the post to be updated
	// Your service layer checks if the user owns the post.
	post, err := app.postService.Update(ctx, postID, input)
	if err != nil {
		var validationErr validator.ValidationError
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		} else if errors.As(err, &validationErr) {
			app.failedValidationResponse(w, r, map[string]any{"error": "Validation failed"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	//send data to client through json marshal
	err = app.writeJSON(w, http.StatusCreated, post)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// deletePostHandler handles DELETE /post/{id} requests
func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	//retrieve post id from request url
	postID := app.readIDParam(r)

	//retrieve userID from request context
	authenticatedUser := r.Context().Value("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//call the service layer that has logic for deleting post from database
	err := app.postService.Delete(ctx, postID, authenticatedUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		} else if err.Error() == "forbidden: user is not the owner of the post" {
			app.badRequestResponse(w, r, err)
			return
		}
		// app.serverErrorResponse(w, r, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a 204 No Content for successful deletions
	w.WriteHeader(http.StatusNoContent)
}
