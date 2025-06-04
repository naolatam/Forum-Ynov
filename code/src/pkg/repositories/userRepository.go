package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func (repository *UserRepository) FindById(id uuid.UUID) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Email, &user.Password, &user.Bio, &user.Avatar, &user.CreatedAt, &user.Role_ID, &user.Google_ID, &user.Github_ID)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}

func (repository *UserRepository) FindByUsernameOrEmail(pseudo string, email string) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE pseudo = ? OR email = ?", pseudo, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Email, &user.Password, &user.Bio, &user.Avatar, &user.CreatedAt, &user.Role_ID, &user.Google_ID, &user.Github_ID)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
