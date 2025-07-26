package service

import (
	"context"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/cliffdoyle/social-network/internal/validator"
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

//Create handles the business logic for creating a new post
func (s *postService) Create(ctx context.Context, input models.PostCreateInput, userID string) (*models.Post, error){
	//Validate the input data from the user
	v:=validator.New()

	models.ValidatePostInput(v,&input)

	if !v.Valid(){
		return nil,models.ErrInvalidFieldInput
	}

	//Map the incoming json input with the actual 'Post' model
	post:=&models.Post{
		
	}

}


func (s *postService) GetByID(ctx context.Context, postID string) (*models.Post, error){

}

func (s *postService) Update(ctx context.Context, postID string, input models.PostUpdateInput) (*models.Post, error){

}

func (s *postService) Delete(ctx context.Context, postID string, userID string) error{

}


