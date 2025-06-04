package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type RoleRepository struct {
	db *sql.DB
}

func (repository *RoleRepository) FindById(id *uuid.UUID) (*models.Role, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var role models.Role
	if rows.Next() {
		err = rows.Scan(&role.ID, &role.Name, &role.Permission)
		if err != nil {
			return nil, err
		}
		return &role, nil
	}
	return nil, errors.New("user not found")
}
