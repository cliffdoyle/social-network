package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

func (app *application) Follow(w http.ResponseWriter, r *http.Request) {
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
    defer cancel()
	me:=r.Context().Value("user_id").(string)
	followeeId:=app.readIDParam(r)
	err:=app.followService.FollowUser(ctx,followeeId,me)
	if err!=nil{
		err = app.writeJSON(w, http.StatusInternalServerError, map[string]any{
		"message": "Failed to exacute follow.Please try again later",
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	}
}

func (app *application) Unfollow(w http.ResponseWriter, r *http.Request) {
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
    defer cancel()
	me:=r.Context().Value("user_id").(string)
	followeeId:=app.readIDParam(r)
	err:=app.followService.UnFollowUser(ctx,followeeId,me)
	if err!=nil{
		err = app.writeJSON(w, http.StatusInternalServerError, map[string]any{
		"message": "Failed to exacute unfollow.Please try again later",
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	}
}

func(app *application)Followers(w http.ResponseWriter, r *http.Request){
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
    defer cancel()
	me:=r.Context().Value("user_id").(string)
    followers,err:=app.followService.ShowFollowers(ctx,me)
	if err!=nil{
			err = app.writeJSON(w, http.StatusInternalServerError, map[string]any{
		"message": "Failed to fetch followers list.Please try again later",
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	}
	json.NewEncoder(w).Encode(map[string]any{
		"followersList":followers,
	})

}

func(app *application)Following(w http.ResponseWriter, r *http.Request){
	ctx,cancel:=context.WithTimeout(context.Background(),3*time.Second)
    defer cancel()
	me:=r.Context().Value("user_id").(string)
    following,err:=app.followService.ShowPeopleIFollow(ctx,me)
	if err!=nil{
			err = app.writeJSON(w, http.StatusInternalServerError, map[string]any{
		"message": "Failed to fetch following list.Please try again later",
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	}
	json.NewEncoder(w).Encode(map[string]any{
		"PeopleIFollow":following,
	})

}