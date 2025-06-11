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

func (repository *CommentRepository) FindById(id *uuid.UUID) (*models.Comment, error) {
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
		err = rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.Post.ID, &comment.Post, &comment.User_ID, &comment.User)
		if err != nil {
			return nil, err
		}
		return &comment, nil
	}
	return nil, errors.New("comment not found")
}

func (repository *CommentRepository) FindByPostID(postId *uuid.UUID) (*[]*models.Comment, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM comments WHERE post_id = ?", postId)
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
