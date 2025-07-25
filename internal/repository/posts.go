package repository

import (
	"github.com/cliffdoyle/social-network/internal/database"
	"github.com/cliffdoyle/social-network/internal/models"
)

// Posts interface defines the interface for posts related database operations
type Posts interface {
	Insert(post *models.Post) error
	Get(id int64) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id int64) error
}

// Define a PostModel struct type which wraps a sql.DB connection pool.
type PostsModel struct {
	DB *database.DB
}

// NewPosts creates a new instance of Posts
func NewPosts(db *database.DB) Posts {
	return &PostsModel{DB: db}
}

// Add a method for inserting a new record in the posts table.
func (m *PostsModel) Insert(post *models.Post) error {
	return nil
}

// Add a method for fetching a specific record from the posts table.
func (m *PostsModel) Get(id int64) (*models.Post, error) {
	return nil, nil
}

// Add a method for updating a specific record in the posts table.
func (m *PostsModel) Update(post *models.Post) error {
	return nil
}

// Add a method for deleting a specific record from the posts table.
func (m *PostsModel) Delete(id int64) error {
	return nil
}
