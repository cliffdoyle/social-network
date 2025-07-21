package models

import (
	"github.com/cliffdoyle/social-network/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

// Set calculates the bcrypt type to manage plaintext and hashed
// passwords
func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

// Matches checks if the provided plaintext password matches
// the stored hashed password.
func (p *Password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case err == bcrypt.ErrMismatchedHashAndPassword:
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

// Check that the Email field is not the empty string, and that it matches the regular
// expression for email addresses that we added in our validator package earlier in the
// book.
func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

// Check if the Password.plaintext field is not nil , then check that the value is not the empty
// string and is between 8 and 72 bytes long.
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "name", "must be provided")
	v.Check(len(user.LastName) <= 500, "name", "must not be more than 500 bytes long")
	// Call the standalone ValidateEmail() helper.
	ValidateEmail(v, user.Email)
	// If the plaintext password is not nil, call the standalone
	// ValidatePasswordPlaintext() helper.
	if user.Password.Plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.Plaintext)
	}
	// If the password hash is ever nil, this will be due to logical error in the codebase
	if user.Password.Hash == nil {
		panic("missing password hash for user")
	}
}
