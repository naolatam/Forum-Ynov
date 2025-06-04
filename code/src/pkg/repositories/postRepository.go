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
	rows, err := repository.db.Query("SELECT * FROM users WHERE id = ?", id)
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
	return nil, errors.New("user not found")
}
