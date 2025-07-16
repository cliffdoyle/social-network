package main

import (
	"encoding/json"
	"net/http"
	"strings"
)
type LoginDetails struct{
Email string `json:"email"`
Password string `json:"password"`
Csrf     string `json:"csrfToken"`

}
func(app *application)LoginHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodPost{
        json.NewEncoder(w).Encode(map[string]any{
			"message":"invalid request",
			"code":http.StatusMethodNotAllowed,
		})
		return
	}

	var loginDetails LoginDetails
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
			"code":http.StatusMethodNotAllowed,
		})
		return
	}

	// //Database call to get user by email

	// user,err:=GetUserByEmail(loginDetails.Email)
	// if err!=nil{
	// 	json.NewEncoder(w).Encode(map[string]any{
	// 		"message":"user does not exist.Please register",
	// 		"code":http.StatusMethodNotAllowed,
	// 	})
	// 	return
	// }


	// //validate password
	// if (!correctPassword){
	// 	json.NewEncoder(w).Encode(map[string]any{
	// 		"message":"user does not exist.Please register",
	// 		"code":http.StatusMethodNotAllowed,
	// 	})
	// 	return
	// }
     
	// app.GenerateSession(w,r,user.USERID)
}
