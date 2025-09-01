package models

import "time"

// ReactionType defines the type of reaction
const (
	ReactionLike    = "like"
	ReactionDislike = "dislike"
)

type Reaction struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    *int      `json:"post_id,omitempty"`
	CommentID *int      `json:"comment_id,omitempty"`
	Type      string    `json:"type"` // like or dislike
	CreatedAt time.Time `json:"created_at"`
}
