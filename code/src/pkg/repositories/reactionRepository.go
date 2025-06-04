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
	rows, err := repository.db.Query("SELECT * FROM users WHERE id = ?", id)
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
	return nil, errors.New("user not found")
}
