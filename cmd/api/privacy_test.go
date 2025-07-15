package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestUpdatePrivacyHandler(t *testing.T) {
	// Setup: create a test database and insert a test user
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (
		uuid TEXT PRIMARY KEY NOT NULL,
		firstName TEXT,
		lastName TEXT,
		nickName TEXT,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		profileStatus BOOLEAN
	)`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users (uuid, firstName, lastName, nickName, email, password, profileStatus)
		VALUES ('test-uuid', 'Test', 'User', 'testuser', 'test@example.com', 'password', 0)`)
	if err != nil {
		t.Fatal(err)
	}

	app := &application{
		db: db,
	}

	// Prepare request body
	body, _ := json.Marshal(map[string]interface{}{
		"uuid":          "test-uuid",
		"profileStatus": true,
	})

	req := httptest.NewRequest(http.MethodPatch, "/update-privacy", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call the handler
	app.updatePrivacyHandler(rr, req)

	// Check response
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check database update
	var status bool
	err = db.QueryRow("SELECT profileStatus FROM users WHERE uuid = ?", "test-uuid").Scan(&status)
	if err != nil {
		t.Fatal(err)
	}
	if !status {
		t.Errorf("expected profileStatus to be true, got false")
	}
}