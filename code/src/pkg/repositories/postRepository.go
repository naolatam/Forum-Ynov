package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func (repository *PostRepository) FindById(id *uuid.UUID) (*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		return &post, nil
	}
	return nil, errors.New("post not found")
}

func (repository *PostRepository) FindByTitle(title *string) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM posts WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	var res = []*models.Post{}
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

func (repository *PostRepository) FindByCategoryId(categoryId *uuid.UUID, limit *int) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 10
	if limit != nil {
		effectiveLimit = *limit
	}

	rows, err := repository.db.Query("SELECT p.* FROM posts p INNER JOIN posts_category c ON p.id = c.post_id WHERE c.category_id = ? LIMIT ?", categoryId, effectiveLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	var res = []*models.Post{}
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}
