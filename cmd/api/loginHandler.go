package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cliffdoyle/social-network/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (app *application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "invalid request",
			"code":    http.StatusMethodNotAllowed,
		})
		return
	}
	fmt.Println("hre55")

	var loginDetails *models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "invalid request",
			"code":    http.StatusInternalServerError,
		})
		return
	}

	email := strings.Trim(loginDetails.Email, "")
	password := strings.Trim(loginDetails.Password, "")
	fmt.Println("email:", email, "password:", password)

	if email == "" || password == "" {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "invalid request",
			"code":    http.StatusBadRequest,
		})
		return
	}

	// //Database call to get user by email

	user, err := app.services.Login(loginDetails)
	if err != nil || user == nil {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "user does not exist.Please register",
			"code":    http.StatusMethodNotAllowed,
		})
		return
	}
	fmt.Println("user:", user)
	ps := user.Password.Hash

	fmt.Println("okay")
	fmt.Println("password details:", ps)
	// valid,err:=ps.Matches(password)
	// if err!=nil ||!valid{
	// 	json.NewEncoder(w).Encode(map[string]any{
	// 		"message":"incorrect credentials.Please try again",
	// 		"code":http.StatusMethodNotAllowed,
	// 	})
	// 	return
	// }
	err = bcrypt.CompareHashAndPassword(ps, []byte(password))
	if err != nil {
		fmt.Println("Error bcrypt", err)
		return
	}

	fmt.Println("okay2")
	app.writeJSON(w,200,map[string]any{
		"success":"loggedin",
	})

	app.GenerateSession(w, *r, user.ID)
}
