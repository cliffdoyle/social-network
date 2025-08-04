package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
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
				Name:    "sessionID",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),
			})
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", session.UserID)
		ctx = context.WithValue(ctx, "session_id", session.SessionID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) GenerateSession(w http.ResponseWriter, r *http.Request, id string) {
	app.logger.Info("--- GENERATE SESSION CALLED for user ID:", id)
	sessionID, err := app.sessionService.PersistSession(id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"message": "failed to save token",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		Secure:   false, 
	}

	http.SetCookie(w, cookie)
	
	app.logger.Info("cookie set",sessionID)

}

func (app *application) LogOut(r *http.Request) {
	sessionID := r.Context().Value("sessionID").(string)
	app.sessionService.DeleteSession(sessionID)
}
