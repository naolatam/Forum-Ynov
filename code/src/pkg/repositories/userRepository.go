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

func (repository *UserRepository) FindByIdOrUsernameOrEmail(id uuid.UUID, pseudo string, email string) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE id = ? OR pseudo = ? OR email = ?", id, pseudo, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
<<<<<<< Updated upstream
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Email, &user.Password, &user.Bio, &user.Avatar, &user.CreatedAt, &user.Role_ID, &user.Google_id, &user.Github_id)
=======
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Email, &user.Password, &user.Bio, &user.Avatar, &user.CreatedAt, &user.Role_ID, &user.Google_ID, &user.Github_ID)
>>>>>>> Stashed changes
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, errors.New("user not found")
}
