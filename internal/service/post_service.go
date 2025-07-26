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
			return nil, models.ErrRecordNotFound
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
			return nil, models.ErrRecordNotFound
		}
		return nil, err
	}

	// for now we just return the post , still waiting for follow and group logic
	return post, nil
}

func (s *postService) Update(ctx context.Context, postID string, input models.PostUpdateInput) (*models.Post, error) {
}

func (s *postService) Delete(ctx context.Context, postID string, userID string) error {
}
