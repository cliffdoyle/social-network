package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)


func (app *application)Authenticator(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
          cookie,err:=r.Cookie("sessionID")
		  if err!=nil ||cookie.Value==""{
			//redirect to login
		  }
		  session,err:=GetSessionFromDB(cookie.Value)
		if session.ExpiresAt().Before(time.Now())||err!=nil{
		http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
			})
		}
		//redirect to login Page
	})
	}
	


func (app *application) GenerateSession(w http.ResponseWriter, r http.Request, id string) {
	session,err := uuid.NewV4()
	if err!=nil{
		json.NewEncoder(w).Encode(map[string]any{
			"message":"Failed to create session",
			"status":http.StatusMethodNotAllowed,
		})
	}
	expiresAt := time.Now().Add(24 * time.Hour)
	sessionID := PersistSession(session) // function that saves the session to the database and returns its session
	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresAt,
		SameSite: http.SameSiteDefaultMode,
	}

	http.SetCookie(w, cookie)
}
