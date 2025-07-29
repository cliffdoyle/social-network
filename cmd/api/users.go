package main

import (
	"net/http"

	"github.com/cliffdoyle/social-network/internal/models"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	var input models.UserRegistrationRequest

	// Parse JSON from request body
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid JSON")
		return
	}

	//Call the service layer to perform the business logic
	user, v, err := app.services.Register(&input)
	if err != nil {
		//Level 500 server error
		app.serverErrorResponse(w, r, err)
	}
	if v != nil && !v.Valid() {
		//Level 422 validation error i.e bad email
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	// Send success response with the created user
	err = app.writeJSON(w, http.StatusCreated, map[string]any{
		"message": "User registered successfully",
		"user":    user,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
