package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cliffdoyle/social-network/internal/database"
	"github.com/cliffdoyle/social-network/internal/models"
)

// Posts interface defines the interface for posts related database operations
type PostRepository interface {
	Insert(ctx context.Context, post *models.Post, audience []string) error
	Get(ctx context.Context, id string) (*models.Post, error)
	Update(ctx context.Context, post *models.Post, newAudience []string) error
	Delete(ctx context.Context, id string) error
}

// Define a PostModel struct type which wraps a sql.DB connection pool.
type PostsModel struct {
	DB *database.DB
}

// NewPosts creates a new instance of Posts
func NewPosts(db *database.DB) PostRepository {
	return &PostsModel{DB: db}
}

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Add a method for inserting a new record in the posts table.
// It can add a post optionally to the post_audience table too
// the operation is performed in a transaction
func (m *PostsModel) Insert(ctx context.Context, post *models.Post, audience []string) error {
	// Begin a new database transaction
	tx, err := m.DB.BeginTx()
	if err != nil {
		return err
	}

	// Defer a rollback in case of an error.
	defer tx.Rollback()

	// Prepare the SQL for inserting into the `posts` table
	query := `
	INSERT INTO posts (id,user_id,group_id,title,content,media_url,media_type,privacy)
	VALUES(?,?,?,?,?,?,?)`

	args := []any{
		post.ID,
		post.UserID,
		post.GroupID,
		post.Title,
		post.Content,
		post.MediaURL,
		post.MediaType,
		post.Privacy,
	}

	// Execute the posts
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	//If the post is private and has an audience, insert into the
	//`post_audience` table.
	if post.Privacy == models.PrivacyPrivate && len(audience) > 0 {
		query2 := `INSERT INTO post_audience (post_id,user_id) VALUES (?,?)`
		stmt, err := tx.PrepareContext(ctx, query2)
		if err != nil {
			return err
		}

		defer stmt.Close()

		for _, userID := range audience {
			_, err := stmt.ExecContext(ctx, post.ID, userID)
			if err != nil {
				return err
			}
		}
	}

	// If everything succeeded, commit the transaction
	return tx.Commit()
}

// Add a method for fetching a specific record from the posts table by its ID.
func (m *PostsModel) Get(ctx context.Context, id string) (*models.Post, error) {
	query := `SELECT id,user_id,group_id,title,content,media_url,media_type,privacy,created_at,updated_at
	FROM posts
	WHERE id = ?`

	var post models.Post

	// Use QueryRowContext on the main connection pool
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.GroupID,
		&post.Title,
		&post.Content,
		&post.MediaURL,
		&post.MediaType,
		&post.Privacy,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &post, nil
}

// Update method for updating a specific record in the posts table and
// post_audience table within a single transaction.
func (m *PostsModel) Update(ctx context.Context, post *models.Post, newAudience []string) error {
	// Begin the transaction
	tx, err := m.DB.BeginTx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// Prepare and execute the UPDATE on the main `posts` table
	query := `UPDATE posts
			SET content = ?, media_url=?, media_type=?,privacy=?, updated_at=?
			WHERE id =?`

	post.UpdatedAt = time.Now()

	args := []any{
		post.Content,
		post.MediaURL,
		post.MediaType,
		post.Privacy,
		post.UpdatedAt,
		post.ID,
	}

	result,err:=tx.ExecContext(ctx,query,args...)
	if err !=nil{
		return  err
	}

	//If no rows were affected we know that no records were affected in the database 
	//therefore none exixted
	rowsAffected,err:=result.RowsAffected()
	if err !=nil{
		return err
	}

	if rowsAffected==0{
		return ErrRecordNotFound
	}

	//Delete all old audience members for this post
	//This runs regardless of the new privacy setting, ensuring we clean up
	//if a post is changed from private to public
	queryAudience:=`DELETE FROM post_audience WHERE post_id=?`

	_,err=tx.ExecContext(ctx,queryAudience,post.ID)
	if err !=nil{
		return err
	}

	//If the new privacy setting is "private", insert the new audience
	if post.Privacy==models.PrivacyPrivate && len(newAudience) > 0{
		query:=`INSERT INTO post_audience (post_id,user_id) VALUES (?,?)`
		stmt,err:=tx.PrepareContext(ctx,query)
		if err !=nil{
			return err
		}
		defer stmt.Close()

		for _,userID:=range newAudience{
			_,err=stmt.ExecContext(ctx,post.ID,userID)
			if err !=nil{
				return err
			}
		}
	}

	//If all steps succeeded, commit the transaction
	return tx.Commit()

	return nil
}

// Add a method for deleting a specific record from the posts table.
// All related rows in `post_audience` will also be deleted automatically
func (m *PostsModel) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM posts WHERE id = ?`

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, we know that the post table didn't contain a record
	// with the provided ID at the moment we tried to delete it. In that case we just
	// return an ErrRecordNotFound error
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
