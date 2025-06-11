package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type ReactionRepository struct {
	db *sql.DB
}

func (repository *ReactionRepository) FindById(id *uuid.UUID) (*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	if rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Post, &reaction.Comment_id, &reaction.Comment, &reaction.User_id, &reaction.User, &reaction.Label)
		if err != nil {
			return nil, err
		}
		return &reaction, nil
	}
	return nil, errors.New("reaction not found")
}

func (repository *ReactionRepository) FindByPostId(postId *uuid.UUID) (*[]*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	var res = []*models.Reaction{}
	for rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Post, &reaction.Comment_id, &reaction.Comment, &reaction.User_id, &reaction.User, &reaction.Label)
		if err != nil {
			return nil, err
		}
		res = append(res, &reaction)
	}
	return &res, nil
}

func (repository *ReactionRepository) FindByCommentId(commentId *uuid.UUID) (*[]*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE comment_id = ?", commentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	var res = []*models.Reaction{}
	for rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Post, &reaction.Comment_id, &reaction.Comment, &reaction.User_id, &reaction.User, &reaction.Label)
		if err != nil {
			return nil, err
		}
		res = append(res, &reaction)
	}
	return &res, nil
}

func (repository *ReactionRepository) FindByUserId(userId *uuid.UUID) (*[]*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	var res = []*models.Reaction{}
	for rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Post, &reaction.Comment_id, &reaction.Comment, &reaction.User_id, &reaction.User, &reaction.Label)
		if err != nil {
			return nil, err
		}
		res = append(res, &reaction)
	}
	return &res, nil
}

func (repository *ReactionRepository) FindByCommentAndUserId(commentId uint32, userId uuid.UUID) (*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE comment_id = ? AND user_id = ?", commentId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	if rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Comment_id, &reaction.User_id, &reaction.Label)
		if err != nil {
			return nil, err
		}
		return &reaction, nil
	}
	return nil, errors.New("reaction not found")
}
func (repository *ReactionRepository) FindByPostAndUserId(postId uint32, userId uuid.UUID) (*models.Reaction, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reactions WHERE post_id = ? AND user_id = ?", postId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reaction models.Reaction
	if rows.Next() {
		err = rows.Scan(&reaction.ID, &reaction.Post_id, &reaction.Comment_id, &reaction.User_id, &reaction.Label)
		if err != nil {
			return nil, err
		}
		return &reaction, nil
	}
	return nil, errors.New("reaction not found")
}

func (repository *ReactionRepository) GetLikeReactionCountOnComment(commentId uint32) (int, error) {
	if repository.db == nil {
		return -1, errors.New("connection to database isn't established")
	}
	var count int
	err := repository.db.QueryRow(`SELECT count(*) FROM reactions 
	WHERE comment_id = ? AND label = 'like'`, commentId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repository *ReactionRepository) GetDislikeReactionCountOnComment(commentId uint32) (int, error) {
	if repository.db == nil {
		return -1, errors.New("connection to database isn't established")
	}
	var count int
	err := repository.db.QueryRow(`SELECT count(*) FROM reactions 
	WHERE comment_id = ? AND label = 'dislike'`, commentId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil

}

func (repository *ReactionRepository) GetLikeReactionCountOnPost(postId uint32) (int, error) {
	if repository.db == nil {
		return -1, errors.New("connection to database isn't established")
	}
	var count int
	err := repository.db.QueryRow(`SELECT count(*) FROM reactions 
	WHERE post_id = ? AND label = 'like'`, postId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (repository *ReactionRepository) GetDislikeReactionCountOnPost(postId uint32) (int, error) {
	if repository.db == nil {
		return -1, errors.New("connection to database isn't established")
	}
	var count int
	err := repository.db.QueryRow(`SELECT count(*) FROM reactions 
	WHERE post_id = ? AND label = 'dislike'`, postId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil

}

func (repository *ReactionRepository) Create(reaction *models.Reaction) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("INSERT INTO reactions ( post_id, comment_id, user_id, label) VALUES ( ?, ?, ?, ?)",
		reaction.Post_id, reaction.Comment_id, reaction.User_id, reaction.Label)
	return err
}

func (repository *ReactionRepository) Update(reaction *models.Reaction) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("UPDATE reactions SET label = ? WHERE id = ?",
		reaction.Label, reaction.ID)
	if err != nil {
		return err
	}
	return nil
}
func (repository *ReactionRepository) Delete(reaction *models.Reaction) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM reactions WHERE id = ?", reaction.ID)
	if err != nil {
		return err
	}
	return nil
}
