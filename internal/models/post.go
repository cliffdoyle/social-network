package models

import (
	"errors"
	"net/url"
	"time"

	"github.com/cliffdoyle/social-network/internal/validator"
)

// PostType defines the type of post, e.g "text" or "link"
type PostPrivacy string

const (
	PrivacyPublic    PostPrivacy = "public"
	PrivacyFollowers PostPrivacy = "followers" // "almost private"
	PrivacyPrivate   PostPrivacy = "private"   // Viewable by a specific list of users
)

var ErrInvalidFieldInput = errors.New("invalid input data. Ensure all fields meet requirement!")

var ErrRecordNotFound = errors.New("record not found")

type Post struct {
	ID        string      `json:"id" db:"id"`
	UserID    string      `json:"userId" db:"user_id"`
	GroupID   string      `json:"groupId,omitzero" db:"group_id"`
	Title     string      `json:"title" db:"title"`
	Content   *string     `json:"content,omitzero" db:"content"`
	MediaURL  *string     `json:"mediaUrl,omitzero" db:"media_url"`
	MediaType *string     `json="mediaUrl,omitzero" db:"media_url"`
	Privacy   PostPrivacy `json:"privacy" db:"privacy"` // "public", "private", "almost_private"
	CreatedAt time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time   `json:"updatedAt" db:"updated_at"`
}

// The struct that maps to the incoming JSON from the user
type PostCreateInput struct {
	GroupID   *string     `json:"groupId"`
	Title     *string     `json:"title"`
	Content   *string     `json:"content"`
	MediaURL  *string     `json:"mediaUrl"`
	MediaType *string     `json:"mediaType"`
	Privacy   PostPrivacy `json:"privacy"`
	Audience  []string    `json:"audience"`
}

// Struct for updates , maps to the incoming JSON request for updates
type PostUpdateInput struct {
	Content   *string     `json:"content"`
	MediaURL  *string     `json:"mediaUrl"`
	MediaType *string     `json:"mediaType"`
	Privacy   PostPrivacy `json:"privacy"`
	Audience  []string    `json:"audience"`
}

// This function validates the core post model. It ensures the data is sane
// and ready for the database, according to the schema.
func ValidatePost(v *validator.Validator, post *Post) {
	// A UserID is always required. This would be set by the server from the session.
	v.Check(post.UserID != "", "userId", "must be provided")

	// A post must have EITHER content OR a media URL, or both. It cannot be completely empty.
	hasContent := post.Content != nil && *post.Content != ""
	hasMedia := post.MediaURL != nil && *post.MediaURL != ""
	v.Check(hasContent || hasMedia, "body", "post must include content or media")

	// If media is provided, its format must be validated.
	if hasMedia {
		_, err := url.ParseRequestURI(*post.MediaURL)
		v.Check(err == nil, "mediaUrl", "must be a valid URL")

		// MediaType must also be provided if there's a URL.
		v.Check(post.MediaType != nil && *post.MediaType != "", "mediaType", "must be provided with mediaUrl")
		if post.MediaType != nil {
			v.Check(validator.PermittedValue(*post.MediaType, "image", "gif"), "mediaType", "must be 'image' or 'gif'")
		}
	} else {
		// If there's no media, there should be no media type.
		v.Check(post.MediaType == nil, "mediaType", "must not be provided without a mediaUrl")
	}

	// Privacy setting must be one of the permitted values.
	v.Check(validator.PermittedValue(post.Privacy, PrivacyPublic, PrivacyFollowers, PrivacyPrivate), "privacy", "must be a valid privacy setting")
}

// This new validation function handles the complete input from the client.
func ValidatePostInput(v *validator.Validator, input *PostCreateInput) {
	// Phase 1: Basic content validation
	// Validate Title: It is required and must have a reasonable length.
	v.Check(input.Title != nil, "title", "must be provided")
	if input.Title != nil {
		v.Check(*input.Title != "", "title", "must not be empty")
		v.Check(len(*input.Title) <= 300, "title", "must not be more than 300 characters long")
	}
	// A post must have EITHER content OR a media URL, or both. It cannot be completely empty.
	hasContent := input.Content != nil && *input.Content != ""
	hasMedia := input.MediaURL != nil && *input.MediaURL != ""
	v.Check(hasContent || hasMedia, "body", "post must include content or media")

	// If media is provided, validate it.
	if hasMedia {
		_, err := url.ParseRequestURI(*input.MediaURL)
		v.Check(err == nil, "mediaUrl", "must be a valid URL")

		v.Check(input.MediaType != nil && *input.MediaType != "", "mediaType", "must be provided with mediaUrl")
		if input.MediaType != nil {
			v.Check(validator.PermittedValue(*input.MediaType, "image", "gif"), "mediaType", "must be 'image' or 'gif'")
		}
	} else {
		v.Check(input.MediaType == nil, "mediaType", "must not be provided without a mediaUrl")
	}

	// Phase 2: Privacy and Audience validation

	// First, validate the privacy setting itself.
	v.Check(validator.PermittedValue(input.Privacy, PrivacyPublic, PrivacyFollowers, PrivacyPrivate), "privacy", "must be a valid privacy setting")

	// Now, apply rules based on the privacy setting.
	switch input.Privacy {
	case PrivacyPublic, PrivacyFollowers:
		// For public or followers-only posts, an audience list MUST NOT be provided.
		// It's an error because it's ambiguous and could lead to bugs.
		v.Check(len(input.Audience) == 0, "audience", "must be empty for public or followers-only posts")

	case PrivacyPrivate:
		// For private posts, an audience list is REQUIRED.
		v.Check(len(input.Audience) > 0, "audience", "must be provided for a private post")

		// You might also want to check for a reasonable limit.
		v.Check(len(input.Audience) <= 100, "audience", "cannot include more than 100 users")
	}

	// Validate GroupID if it is provided.
	// This is optional, so we only check it if the pointer is not nil.
	if input.GroupID != nil {
		// Here you might check if the GroupID looks like a valid UUID, or if the
		// group actually exists in the database (this check is often done in the service layer).
		v.Check(*input.GroupID != "", "groupId", "must not be an empty string if provided")
	}
}
