package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"errors"

	"github.com/google/uuid"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
}

// FindByCommentAndUserId retrieves a reaction by comment ID and user ID.
func (s *ReactionService) FindByCommentAndUserId(commentId uint32, userId uuid.UUID) *models.Reaction {
	if userId == uuid.Nil {
		return nil // Invalid user ID, do not search for reaction
	}
	if commentId == 0 {
		return nil // Invalid comment ID, do not search for reaction
	}
	reaction, err := s.repo.FindByCommentAndUserId(commentId, userId)
	if err != nil {
		return nil // Error occurred while searching for reaction
	}
	return reaction
}

// FindByPostAndUserId retrieves a reaction by post ID and user ID.
func (s *ReactionService) FindByPostAndUserId(postId uint32, userId uuid.UUID) *models.Reaction {
	if userId == uuid.Nil {
		return nil // Invalid user ID, do not search for reaction
	}
	if postId == 0 {
		return nil // Invalid comment ID, do not search for reaction
	}
	reaction, err := s.repo.FindByPostAndUserId(postId, userId)
	if err != nil {
		return nil // Error occurred while searching for reaction
	}
	return reaction
}

// GetLikeReactionCountOnComment retrieves the count of "like" reactions on a comment.
func (s *ReactionService) GetLikeReactionCountOnComment(comment *models.Comment) int {
	if comment == nil {
		return 0
	}
	reactionCount, _ := s.repo.GetLikeReactionCountOnComment(comment.ID)

	return reactionCount
}

// GetDislikeReactionCountOnComment retrieves the count of "dislike" reactions on a comment.
func (s *ReactionService) GetDislikeReactionCountOnComment(comment *models.Comment) int {
	if comment == nil {
		return 0
	}
	reactionCount, _ := s.repo.GetDislikeReactionCountOnComment(comment.ID)

	return reactionCount
}

// GetLikeReactionCountOnPost retrieves the count of "like" reactions on a post.
func (s *ReactionService) GetLikeReactionCountOnPost(post *models.Post) int {
	if post == nil {
		return 0
	}
	reactionCount, _ := s.repo.GetLikeReactionCountOnPost(post.ID)

	return reactionCount
}

// GetDislikeReactionCountOnPost retrieves the count of "dislike" reactions on a post.
func (s *ReactionService) GetDislikeReactionCountOnPost(post *models.Post) int {
	if post == nil {
		return 0
	}
	reactionCount, _ := s.repo.GetDislikeReactionCountOnPost(post.ID)

	return reactionCount
}

// Create creates a new reaction for a post or comment by a user with a specific label.
func (s *ReactionService) Create(postId *uint32, commentId *uint32, user *models.User, label string) error {
	if label != "like" && label != "dislike" {
		return errors.New("invalid label") // Invalid label, do not create reaction
	}
	reaction := &models.Reaction{
		Post_id:    postId,
		Comment_id: commentId,
		User_id:    user.ID,
		Label:      label,
	}

	return s.repo.Create(reaction)
}

// Update updates an existing reaction with a new label.
func (s *ReactionService) Update(reaction *models.Reaction) bool {
	if reaction == nil {
		return false // Reaction is nil, do not update
	}
	if reaction.Label != "like" && reaction.Label != "dislike" {
		return false // Invalid label, do not update reaction
	}

	return s.repo.Update(reaction) == nil
}

// Delete removes a reaction from the repository.
func (s *ReactionService) Delete(reaction *models.Reaction) bool {
	if reaction == nil {
		return false // Reaction is nil, do not delete
	}
	if reaction.Label != "like" && reaction.Label != "dislike" {
		return false // Invalid label, do not delete reaction
	}

	return s.repo.Delete(reaction) == nil
}
