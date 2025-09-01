package repository

import (
	"database/sql"
	"errors"

	"github.com/cliffdoyle/social-network/internal/models"
)

type ReactionRepository struct {
	DB *sql.DB
}

func NewReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{DB: db}
}

func (r *ReactionRepository) AddReaction(reaction *models.Reaction) error {
	query := `INSERT INTO reactions (user_id, post_id, comment_id, type) VALUES (?, ?, ?, ?)`
	_, err := r.DB.Exec(query, reaction.UserID, reaction.PostID, reaction.CommentID, reaction.Type)
	return err
}

func (r *ReactionRepository) RemoveReaction(userID int, postID, commentID *int) error {
	query := `DELETE FROM reactions WHERE user_id = ? AND post_id IS ? AND comment_id IS ?`
	_, err := r.DB.Exec(query, userID, postID, commentID)
	return err
}

func (r *ReactionRepository) GetReaction(userID int, postID, commentID *int) (*models.Reaction, error) {
	query := `SELECT id, user_id, post_id, comment_id, type, created_at FROM reactions WHERE user_id = ? AND post_id IS ? AND comment_id IS ?`
	row := r.DB.QueryRow(query, userID, postID, commentID)
	var reaction models.Reaction
	err := row.Scan(&reaction.ID, &reaction.UserID, &reaction.PostID, &reaction.CommentID, &reaction.Type, &reaction.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &reaction, nil
}

func (r *ReactionRepository) CountReactions(postID, commentID *int, reactionType string) (int, error) {
	query := `SELECT COUNT(*) FROM reactions WHERE post_id IS ? AND comment_id IS ? AND type = ?`
	row := r.DB.QueryRow(query, postID, commentID, reactionType)
	var count int
	err := row.Scan(&count)
	return count, err
}
