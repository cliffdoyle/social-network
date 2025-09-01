package service

import (
	"errors"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
)

type ReactionService struct {
	repo *repository.ReactionRepository
}

func NewReactionService(repo *repository.ReactionRepository) *ReactionService {
	return &ReactionService{repo: repo}
}

// React allows a user to like or dislike a post or comment
func (s *ReactionService) React(userID int, postID, commentID *int, reactionType string) error {
	if reactionType != models.ReactionLike && reactionType != models.ReactionDislike {
		return errors.New("invalid reaction type")
	}

	existing, err := s.repo.GetReaction(userID, postID, commentID)
	if err != nil {
		return err
	}

	if existing != nil {
		if existing.Type == reactionType {
			// Already reacted with the same type, remove reaction (toggle off)
			return s.repo.RemoveReaction(userID, postID, commentID)
		}
		// Update reaction type
		err := s.repo.RemoveReaction(userID, postID, commentID)
		if err != nil {
			return err
		}
	}

	reaction := &models.Reaction{
		UserID:    userID,
		PostID:    postID,
		CommentID: commentID,
		Type:      reactionType,
	}
	return s.repo.AddReaction(reaction)
}

func (s *ReactionService) CountReactions(postID, commentID *int, reactionType string) (int, error) {
	return s.repo.CountReactions(postID, commentID, reactionType)
}
