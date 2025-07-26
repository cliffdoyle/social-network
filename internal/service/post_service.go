package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/cliffdoyle/social-network/internal/validator"
	"github.com/google/uuid"
)

// PostSErvice defines the interface for the post business logic
// Its methods will be called by the HTTP handlers
type PostService interface {
	Create(ctx context.Context, input models.PostCreateInput, userID string) (*models.Post, error)
	GetByID(ctx context.Context, postID string) (*models.Post, error)
	Update(ctx context.Context, postID string, input models.PostUpdateInput) (*models.Post, error)
	Delete(ctx context.Context, postID string, userID string) error
}

// postService struct implements the UserService interface
type postService struct {
	repo repository.PostRepository
}

// NewPostService creates a new instance of the postService
func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

// Create handles the business logic for creating a new post
func (s *postService) Create(ctx context.Context, input models.PostCreateInput, userID string) (*models.Post, error) {
	// Validate the input data from the user
	v := validator.New()

	models.ValidatePostInput(v, &input)

	if !v.Valid() {
		return nil, models.ErrInvalidFieldInput
	}

	// Map the incoming json input with the actual 'Post' model
	post := &models.Post{
		ID:        uuid.NewString(), // Generates new unique ID for the post
		UserID:    userID,
		GroupID:   *input.GroupID,
		Title:     *input.Title,
		Content:   input.Content,
		MediaURL:  input.MediaURL,
		Privacy:   input.Privacy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Call the repository to insert the post into the database.
	// The repository will handle the transaction for `posts` and `post_audience`
	err := s.repo.Insert(ctx, post, input.Audience)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err//the handler to inspect and decide which HTTP status code to return 
		}
		return nil, err
	}

	// Return the newly created post object
	return post, nil
}

// GetByID handles fetching a single post
func (s *postService) GetByID(ctx context.Context, postID string) (*models.Post, error) {
	post, err := s.repo.Get(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	// for now we just return the post , still waiting for follow and group logic
	return post, nil
}

func (s *postService) Update(ctx context.Context, postID string, input models.PostUpdateInput) (*models.Post, error) {
	// Fetch the existing post to ensure it exists and to get its owner ID
	existsPost, err := s.repo.Get(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	// Conditionally apply updates only for fields that were provided in the input
	if input.Content != nil {
		existsPost.Content = input.Content
	}

	if input.MediaURL != nil {
		existsPost.MediaURL = input.MediaURL

		existsPost.MediaType = input.MediaType
	}

	if input.Privacy != "" {
		existsPost.Privacy = input.Privacy
	}

	var newAudience []string = nil

	if input.Audience != nil {
		newAudience = input.Audience
	}

	v := validator.New()

	// Pass the raw input DTO from user for validation before final merge
	models.ValidatePostUpdateInput(v, &input, existsPost.Privacy)

	// After validating the input we validate the final post object
	// ensures the update doesn't result in an empty post
	models.ValidatePost(v, existsPost)

	if !v.Valid() {
		return nil,&validator.ValidationError{Errors: v.Errors}
	}

	// Call the repository
	err = s.repo.Update(ctx, existsPost, newAudience)
	if err != nil {
		return nil, err
	}
	return existsPost, nil
}

func (s *postService) Delete(ctx context.Context, postID string, userID string) error {
}
