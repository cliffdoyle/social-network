package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cliffdoyle/social-network/internal/models"
)
func(app *application)LoginHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodPost{
		json.NewEncoder(w).Encode(map[string]any{
			"message":"invalid request",
			"code":http.StatusMethodNotAllowed,
		})
		return
	}

	var loginDetails *models.LoginRequest
	err:=json.NewDecoder(r.Body).Decode(&loginDetails)
	if err!=nil{
		 json.NewEncoder(w).Encode(map[string]any{
			"message":"invalid request",
			"code":http.StatusInternalServerError,
		})
		return
	}
 
email:=strings.Trim(loginDetails.Email,"")
password:=strings.Trim(loginDetails.Password,"")

	if email==""||password==""{
		json.NewEncoder(w).Encode(map[string]any{
			"message":"invalid request",
			"code":http.StatusBadRequest,
		})
		return
	}

	// //Database call to get user by email

	user,err:=app.services.Login(loginDetails)
	if err!=nil || user==nil{
		json.NewEncoder(w).Encode(map[string]any{
			"message":"user does not exist.Please register",
			"code":http.StatusMethodNotAllowed,
		})
		return
	}
     var p *models.Password
	valid,err:=p.Matches(password)
	if err!=nil ||!valid{
		json.NewEncoder(w).Encode(map[string]any{
			"message":"incorrect credentials.Please try again",
			"code":http.StatusMethodNotAllowed,
		})
		return
	}
     
	app.GenerateSession(w,*r,user.ID)
}
