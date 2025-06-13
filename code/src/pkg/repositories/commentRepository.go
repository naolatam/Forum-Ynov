package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type CommentRepository struct {
	db *sql.DB
}

// FindById retrieves a comment by its ID.
func (repository *CommentRepository) FindById(id uint32) (*models.Comment, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM comments WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comment models.Comment
	if rows.Next() {
		err = rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.Post_id, &comment.User_ID)
		if err != nil {
			return nil, err
		}
		comment.Post.ID = comment.Post_id
		comment.User.ID = comment.User_ID
		return &comment, nil
	}
	return nil, errors.New("comment not found")
}

// FindByPostId retrieves all comments associated with a specific post ID.
func (repository *CommentRepository) FindByPostId(postId uint32) (*[]*models.Comment, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM comments WHERE post_id = ? ORDER BY createdAt DESC", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Comment{}
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.Post.ID, &comment.User_ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &comment)
	}
	return &res, nil
}

// FindByUserId retrieves all comments made by a specific user.
func (repository *CommentRepository) FindByUserId(userId *uuid.UUID) (*[]*models.Comment, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM comments WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comment models.Comment
	var res = []*models.Comment{}
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.Post.ID, &comment.Post, &comment.User_ID, &comment.User)
		if err != nil {
			return nil, err
		}
		res = append(res, &comment)
	}
	return &res, nil
}

// GetUserCommentCount retrieves the count of comments made by a specific user.
func (repository *CommentRepository) GetUserCommentCount(userId *uuid.UUID) (int, error) {
	if repository.db == nil {
		return -1, errors.New("unable to connect to database")
	}
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM comments WHERE user_id = ?", userId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// Create inserts a new comment into the database.
func (repository *CommentRepository) Create(comment *models.Comment) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("INSERT INTO comments ( content, createdAt, post_id, user_id) VALUES ( ?, ?, ?, ?)",
		comment.Content, comment.CreatedAt, comment.Post_id, comment.User_ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a comment from the database.
func (repository *CommentRepository) Delete(comment *models.Comment) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM comments WHERE id = ?", comment.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing comment in the database.
func (repository *CommentRepository) Update(comment *models.Comment) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("UPDATE comments SET content = ?, post_id = ?, user_id = ? WHERE id = ?",
		comment.Content, comment.Post_id, comment.User_ID, comment.ID)
	if err != nil {
		return err
	}
	return nil
}
