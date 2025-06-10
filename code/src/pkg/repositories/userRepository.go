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

func (repository *UserRepository) FindByUsernameOrEmail(pseudo *string, email *string) (*models.User, error) {
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

func (repository *UserRepository) FindByEmail(email *string) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE email = ?", email)
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

func (repository *UserRepository) FindByUsername(username *string) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE pseudo = ?", username)
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

func (repository *UserRepository) FindByGoogleID(googleID string) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE google_id = ?", googleID)
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

func (repository *UserRepository) FindByGithubID(githubID int64) (*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE github_id = ?", githubID)
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

func (repository *UserRepository) Update(user *models.User) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}

	_, err := repository.db.Exec("UPDATE users SET pseudo = ?, email = ?, password = ?, bio = ?, avatar = ?, role_id = ?, google_id = ?, github_id = ? WHERE id = ?",
		user.Pseudo, user.Email, user.Password, user.Bio, user.Avatar, user.Role_ID, user.Google_ID, user.Github_ID, user.ID)
	if err != nil {
		return err
	}
	return nil

}

func (repository *UserRepository) Create(user *models.User) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	if user == nil {
		return errors.New("user cannot be nil")
	}

	_, err := repository.db.Exec("INSERT INTO users (id, pseudo, email, password, bio, avatar, createdAt, role_id, google_id, github_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.Pseudo, user.Email, user.Password, user.Bio, user.Avatar, user.CreatedAt, user.Role_ID, user.Google_ID, user.Github_ID)
	if err != nil {
		return err
	}
	return nil

}

func (repository *UserRepository) GetUserCount() int {
	if repository.db == nil {
		return -1
	}

	var userCount int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
	if err != nil {
		return -1
	}
<<<<<<< Updated upstream
	return userCount
}
=======
<<<<<<< Updated upstream
	return nil
}
=======
	return userCount
}

func (repository *UserRepository) Delete(userId *uuid.UUID) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	if userId == nil || *userId == uuid.Nil {
		return errors.New("user ID cannot be nil")
	}

	_, err := repository.db.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		return err
	}
	return nil
}
>>>>>>> Stashed changes
>>>>>>> Stashed changes
