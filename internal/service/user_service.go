package service

import (
	"time"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/cliffdoyle/social-network/internal/validator"
)

// UserService interface defines the methods that a user service should implement
// for user-related business logic
type UserService interface {
	Register(input *models.UserRegistrationRequest) (*models.User, *validator.Validator, error)
	
}

// userService struct implements the UserService interface
type userService struct {
	repo repository.UserRepository
	
}

// NewUserService creates a new instance of the userService
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Register handles the business logic for user registratiobn
func (s *userService) Register(input *models.UserRegistrationRequest) (*models.User, *validator.Validator, error) {
	// Parse the date of birth string into time.Time
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		v := validator.New()
		v.AddError("dateOfBirth", "must be a valid date in YYYY-MM-DD format")
	}

	// Create the main user model
	user := &models.User{
		Email:         input.Email,
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

	//Set and hash the password using the model's method
	err = user.Password.Set(input.Password)
	if err != nil {
		//server error during hashing
		return nil, nil, err
	}

	//Validate the user struct using the model's validation function
	v := validator.New()

	//replaces all the v.checks*
	models.ValidateUser(v, user)

	//Check for a duplicate email *after* initial validation passes
	if v.Valid() {
		existingUser, err := s.repo.FindByEmail(input.Email)
		if err != nil {
			//database error occurred
			return nil, nil, err
		}

		if existingUser != nil {
			v.AddError("email", repository.ErrDuplicateEmail.Error())
		}
	}

	if !v.Valid() {
		return nil, v, nil
	}

	//Save the user to the database
	err = s.repo.Insert(user)
	if err != nil {
		return nil, nil, err
	}

	return user, nil, nil

}
