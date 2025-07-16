package main

import (
	"encoding/json"
	"net/http"
)

// Request body structure
type privacyUpdateRequest struct {
	UUID          string `json:"uuid"` // or use your authentication/session to get user
	ProfileStatus bool   `json:"profileStatus"`
}

func (app *application) updatePrivacyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch && r.Method != http.MethodPost {
		app.methodNotAllowedResponse(w, r)
		return
	}

	var req privacyUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Update the user's profileStatus in the database
	_, err = app.db.Exec(
		"UPDATE users SET profileStatus = ? WHERE uuid = ?",
		req.ProfileStatus, req.UUID,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	data := map[string]any{
		"message": "Privacy updated successfully",
	}
	app.writeJSON(w, http.StatusOK, data)
}
