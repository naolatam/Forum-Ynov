package repositories

import (
	"Forum-back/internal/config"
	"Forum-back/pkg/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func (repository *UserRepository) Init() bool {
	var err error
	repository.db, err = config.OpenDBConnection()
	statusInit := true
	if err != nil {
		log.Println("[UserRepository] Failed to connect to database")
		statusInit = false
	}
	return statusInit
}

func (repository *UserRepository) Close() {
	err := repository.db.Close()
	if err != nil {
		log.Println("[UserRepository] Failed to close the connection to database")
	}
}

func (repository *UserRepository) FindById(id uuid.UUID) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("Connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Email, &user.Password, &user.Bio, &user.Avatar, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("User not found")
}
