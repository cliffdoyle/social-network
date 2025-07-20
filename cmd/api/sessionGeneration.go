package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func (app *application) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionID")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		session, err := app.sessionService.GetSessionFromDB(cookie.Value)
		if session.Expires.Before(time.Now()) || err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:    "session_id",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
			})
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	})
}

func (app *application) GenerateSession(w http.ResponseWriter, r http.Request, id string) {
	sessionID := uuid.NewV4()
	sessionID, err := app.sessionService.PersistSession(id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "failed to save token",
			"status":  http.StatusInternalServerError,
		}) // function that saves the session to the database and returns its session
		cookie := &http.Cookie{
			Name:     "sessionID",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
			SameSite: http.SameSiteDefaultMode,
		}

		http.SetCookie(w, cookie)
	}
}
