package repository

import (
	"database/sql"
	"errors"

	"github.com/cliffdoyle/social-network/internal/database"
	"github.com/cliffdoyle/social-network/internal/models"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email: user with this email already exists")
)

// UserRepository interface defines the interface for user related database operations
type UserRepository interface {
	Insert(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

// UserRepository struct implements the UserRepository interface
type userRepository struct {
	DB *database.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{DB: db}
}

// FindByEmail checks if a user with a given email eists
// It returns the user if found, or sql.ErrNoRows if not found
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id,password_hash FROM users WHERE email =?"

	err := r.DB.QueryRow(query, email).Scan(&user.ID,&user.Password.Hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//No user found, not an application error
			return nil, nil
		}
		//A real database error occurred
		return nil, err
	}
	return &user, nil
}

// Insert adds a new user to the database
func (r *userRepository) Insert(user *models.User) error {
	insertQuery := `
        INSERT INTO users (email, password_hash, first_name, last_name, date_of_birth, nickname, about_me, is_private, email_verified, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING id`

	// Execute the query and get the generated ID
	err := r.DB.QueryRow(insertQuery,
		user.Email,
		user.Password.Hash,
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
	return err
}
