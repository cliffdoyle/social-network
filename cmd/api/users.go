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

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, "Invalid JSON")
        return
    }

    // Validate input
    v := validator.New()
    
    // Basic validation
    v.Check(input.Email != "", "email", "must be provided")
    v.Check(validator.Matches(input.Email, validator.EmailRX), "email", "must be a valid email address")
    v.Check(input.Password != "", "password", "must be provided")
    v.Check(len(input.Password) >= 6, "password", "must be at least 6 characters long")
    v.Check(input.FirstName != "", "firstName", "must be provided")
    v.Check(input.LastName != "", "lastName", "must be provided")
    v.Check(input.DateOfBirth != "", "dateOfBirth", "must be provided")

    // Parse date of birth
    dob, err := time.Parse("2006-01-02", input.DateOfBirth)
    if err != nil {
        v.AddError("dateOfBirth", "must be a valid date in YYYY-MM-DD format")
    }

    if !v.Valid() {
        app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    // Create user
    user := &models.User{
        Email:        input.Email,
        PasswordHash: string(hashedPassword),
        FirstName:    input.FirstName,
        LastName:     input.LastName,
        DateOfBirth:  dob,
        IsPrivate:    false,
        EmailVerified: false,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    // Set optional fields
    if input.Nickname != "" {
        user.Nickname = &input.Nickname
    }
    if input.AboutMe != "" {
        user.AboutMe = &input.AboutMe
    }
    
    // Remove sensitive data before response
    user.PasswordHash = ""
    
    err = app.writeJSON(w, http.StatusCreated, map[string]any{
        "message": "User registered successfully",
        "user":    user,
    })
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }
}