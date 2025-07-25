package models

import (
	"net/url"
	"time"

	"github.com/cliffdoyle/social-network/internal/validator"
)

// PostType defines the type of post, e.g "text" or "link"
type PostType string

const (
	PostTypeText   PostType = "text"
	PostTypeLink   PostType = "link"
	PostTypeImage  PostType = "image"
	PostTypeVideos PostType = "videos"
)

type Post struct {
	ID      string   `json:"id" db:"id"`
	UserID  string   `json:"userId" db:"user_id"`
	GroupID string   `json:"groupId" db:"group_id"`
	Type    PostType `json:"type" db:"type"`
	Title   string   `json:"title" db:"title"`
	Content *string  `json:"content" db:"content"`
	URL     *string  `json:"url,omitempty" db:"url"`
	// ImagePath *string   `json:"imagePath,omitempty" db:"image_path"`
	Privacy   string    `json:"privacy" db:"privacy"` // "public", "private", "almost_private"
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func ValidatePost(v *validator.Validator, post *Post) {
	// Basic field validation
	v.Check(post.UserID != "", "userId", "must be provided")
	v.Check(post.GroupID != "", "groupId", "must be provided")
	v.Check(post.Title != "", "title", "must be provided")
	v.Check(len(post.Title) <= 300, "title", "must not be more than 300 characters long")

	// 1. Update the list of permitted types
	v.Check(
		validator.PermittedValue(post.Type, PostTypeText, PostTypeLink, PostTypeImage, PostTypeVideos),
		"type",
		"must be a valid post type (text, link, image, or videos)",
	)

	v.Check(validator.PermittedValue(post.Privacy, "public", "private", "almost_private"), "privacy", "must be a valid privacy setting")

	// 2. Update the conditional logic with grouped cases
	switch post.Type {
	case PostTypeText:
		// A text post must have content and must not have a URL.
		v.Check(post.Content != nil && *post.Content != "", "content", "must be provided for a text post")
		v.Check(post.URL == nil, "url", "must be empty for a text post")

	case PostTypeLink, PostTypeImage, PostTypeVideos:
		// A link, image, or video post must have a URL and must not have content.
		v.Check(post.URL != nil && *post.URL != "", "url", "must be provided for this post type")

		// Validate that the URL is a valid format.
		if post.URL != nil && *post.URL != "" {
			_, err := url.ParseRequestURI(*post.URL)
			v.Check(err == nil, "url", "must be a valid URL")
		}

		v.Check(post.Content == nil, "content", "must be empty for this post type")
	}

}
