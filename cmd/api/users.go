package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/cliffdoyle/social-network/internal/models"
    "github.com/cliffdoyle/social-network/internal/validator"
    "golang.org/x/crypto/bcrypt"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        app.methodNotAllowedResponse(w, r)
        return
    }

    var input models.UserRegistrationRequest

    // Parse JSON from request body
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, "Invalid JSON")
        return
    }

    // Validate the input using your existing validator
    v := validator.New()
    
    v.Check(input.Email != "", "email", "must be provided")
    v.Check(validator.Matches(input.Email, validator.EmailRX), "email", "must be a valid email address")
    v.Check(input.Password != "", "password", "must be provided")
    v.Check(len(input.Password) >= 6, "password", "must be at least 6 characters long")
    v.Check(input.FirstName != "", "firstName", "must be provided")
    v.Check(input.LastName != "", "lastName", "must be provided")
    v.Check(input.DateOfBirth != "", "dateOfBirth", "must be provided")

    // Parse the date of birth string into time.Time
    dob, err := time.Parse("2006-01-02", input.DateOfBirth)
    if err != nil {
        v.AddError("dateOfBirth", "must be a valid date in YYYY-MM-DD format")
    }

    // If validation failed, return errors
    if !v.Valid() {
        app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
        return
    }

    // Check if user with this email already exists
    var existingUserID string
    checkQuery := "SELECT id FROM users WHERE email = ?"
    err = app.db.QueryRow(checkQuery, input.Email).Scan(&existingUserID)
    if err == nil {
        // User exists, return error
        v.AddError("email", "a user with this email address already exists")
        app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
        return
    }

    // Hash the password using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // Create the user struct with the validated data
    user := &models.User{
        Email:         input.Email,
        PasswordHash:  string(hashedPassword),
        FirstName:     input.FirstName,
        LastName:      input.LastName,
        DateOfBirth:   dob,
        IsPrivate:     false,
        EmailVerified: false,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    // Handle optional fields (nickname and aboutMe)
    if input.Nickname != "" {
        user.Nickname = &input.Nickname
    }
    if input.AboutMe != "" {
        user.AboutMe = &input.AboutMe
    }

    // Insert the user into the database
    insertQuery := `
        INSERT INTO users (email, password_hash, first_name, last_name, date_of_birth, nickname, about_me, is_private, email_verified, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

    // Execute the query and get the generated ID
    err = app.db.QueryRow(insertQuery,
        user.Email,
        user.PasswordHash,
        user.FirstName,
        user.LastName,
        user.DateOfBirth,
        user.Nickname,
        user.AboutMe,
        user.IsPrivate,
        user.EmailVerified,
        user.CreatedAt,
        user.UpdatedAt,
    ).Scan(&user.ID)

    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // Remove sensitive data before sending response
    user.PasswordHash = ""
    
    // Send success response with the created user
    err = app.writeJSON(w, http.StatusCreated, map[string]any{
        "message": "User registered successfully",
        "user":    user,
    })
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }
}
