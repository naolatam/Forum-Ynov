package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *sql.DB
}

func (repository *CategoryRepository) FindById(id *uuid.UUID) (*models.Category, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category models.Category
	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}
	return nil, errors.New("user not found")
}
