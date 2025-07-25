package models

import (
	"time"
)

// User represents a user in the social network
type User struct {
	ID            string    `json:"id" db:"id"`
	Email         string    `json:"email" db:"email"`
	Password      Password  `json:"Hash" db:"password_hash"` // Never include in JSON responses
	FirstName     string    `json:"firstName" db:"first_name"`
	LastName      string    `json:"lastName" db:"last_name"`
	DateOfBirth   time.Time `json:"dateOfBirth" db:"date_of_birth"`
	Nickname      *string   `json:"nickname,omitempty" db:"nickname"`      // Optional field
	AboutMe       *string   `json:"aboutMe,omitempty" db:"about_me"`       // Optional field
	AvatarPath    *string   `json:"avatarPath,omitempty" db:"avatar_path"` // Optional field
	IsPrivate     bool      `json:"isPrivate" db:"is_private"`
	EmailVerified bool      `json:"emailVerified" db:"email_verified"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

// UserRegistrationRequest represents the data needed to register a new user
type UserRegistrationRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"` // Will be parsed to time.Time
	Nickname    string `json:"nickname,omitempty"`
	AboutMe     string `json:"aboutMe,omitempty"`
}

// Custom password type  is a struct containing the plaintext and hashed
// versions of the password for a user. The plaintext field is a *pointer* to a string,
// so that we're able to distinguish between a plaintext password not being present in
// the struct at all, versus a plaintext password which is the empty string "".
type Password struct {
	Plaintext *string
	Hash      []byte
}

// UserUpdateRequest represents the data that can be updated for a user
type UserUpdateRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Nickname  *string `json:"nickname,omitempty"`
	AboutMe   *string `json:"aboutMe,omitempty"`
	IsPrivate *bool   `json:"isPrivate,omitempty"`
}

// // UserPublicProfile represents the public view of a user (without sensitive data)
// type UserPublicProfile struct {
// 	ID          string    `json:"id"`
// 	FirstName   string    `json:"firstName"`
// 	LastName    string    `json:"lastName"`
// 	Nickname    *string   `json:"nickname,omitempty"`
// 	AboutMe     *string   `json:"aboutMe,omitempty"`
// 	AvatarPath  *string   `json:"avatarPath,omitempty"`
// 	IsPrivate   bool      `json:"isPrivate"`
// 	CreatedAt   time.Time `json:"createdAt"`
// }



// Post represents a post in the social network


// // Group represents a group in the social network
// type Group struct {
// 	ID          string    `json:"id" db:"id"`
// 	Name        string    `json:"name" db:"name"`
// 	Description string    `json:"description" db:"description"`
// 	CreatorID   string    `json:"creatorId" db:"creator_id"`
// 	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
// 	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
// }

// // Event represents an event in a group
// type Event struct {
// 	ID          string    `json:"id" db:"id"`
// 	GroupID     string    `json:"groupId" db:"group_id"`
// 	CreatorID   string    `json:"creatorId" db:"creator_id"`
// 	Title       string    `json:"title" db:"title"`
// 	Description string    `json:"description" db:"description"`
// 	DateTime    time.Time `json:"dateTime" db:"date_time"`
// 	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
// 	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
// }

// // Chat represents a chat message
// type Chat struct {
// 	ID         string    `json:"id" db:"id"`
// 	SenderID   string    `json:"senderId" db:"sender_id"`
// 	ReceiverID *string   `json:"receiverId,omitempty" db:"receiver_id"` // nil for group chats
// 	GroupID    *string   `json:"groupId,omitempty" db:"group_id"`       // nil for private chats
// 	Message    string    `json:"message" db:"message"`
// 	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
// }

// // Notification represents a notification
// type Notification struct {
// 	ID             string    `json:"id" db:"id"`
// 	UserID         string    `json:"userId" db:"user_id"`
// 	Type           string    `json:"type" db:"type"` // "follow_request", "group_invitation", etc.
// 	SenderID       string    `json:"senderId" db:"sender_id"`
// 	RelatedID      *string   `json:"relatedId,omitempty" db:"related_id"` // group_id, post_id, etc.
// 	Message        string    `json:"message" db:"message"`
// 	IsRead         bool      `json:"isRead" db:"is_read"`
// 	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
// }

// type Comments struct{
// 	CommUID string `json:"commuid"`
// 	Content string `json:"content"`
// 	Sender  string `json:"sender"`
// }

// type Reaction struct{
// 	Likes int
// 	Dislikes int

// }

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type Sessions struct {
	SessionID string    `json:"sessionID"`
	UserID    string    `json:"userID"`
	Expires   time.Time `json:"expiryTime"`
}
