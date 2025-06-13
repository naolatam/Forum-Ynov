package repositories

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/utils"
	"database/sql"
	"errors"
	"html/template"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

// FindById retrieves a user by their ID.
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
		user.AvatarBase64 = template.URL(utils.ConvertBytesToBase64(user.Avatar, "image/png"))
		return &user, nil
	}
	return nil, errors.New("user not found")
}

// FindMultipleByAny retrieves users based on a query that matches pseudo, email, or ID.
func (repository *UserRepository) FindMultipleByAny(query string) (*[]*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}

	rows, err :=
		repository.db.Query(`SELECT u.id, u.pseudo, u.avatar, r.id, r.name, r.permission FROM users u
		INNER JOIN roles r ON u.role_id = r.id
		WHERE u.pseudo LIKE ? OR u.email LIKE ? OR u.id LIKE ?;`, query, query, query)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		var role models.Role
		err = rows.Scan(&user.ID, &user.Pseudo, &user.Avatar, &role.ID, &role.Name, &role.Permission)
		user.Role = role
		user.Role_ID = role.ID
		user.AvatarBase64 = template.URL(utils.ConvertBytesToBase64(user.Avatar, "image/png"))

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &users, nil
}

// FindByUsernameOrEmail retrieves a user by their username or email.
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

// FindByEmail retrieves a user by their email address.
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

// FindByUsername retrieves a user by their username.
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

// FindByGoogleID retrieves a user by their Google ID.
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

// FindByGithubID retrieves a user by their GitHub ID.
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

// Update modifies an existing user in the database.
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

// Create inserts a new user into the database.
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

// GetAllUsers retrieves all users from the database.
func (repository *UserRepository) GetAllUsers() ([]*models.User, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query(`
		SELECT u.id, u.pseudo, u.createdAt, r.id, r.name, r.permission
		FROM users u
		INNER JOIN roles r ON u.role_id = r.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		var role models.Role
		err := rows.Scan(
			&user.ID, &user.Pseudo, &user.CreatedAt, &role.ID, &role.Name, &role.Permission,
		)
		if err != nil {
			return nil, err
		}
		user.Role = role
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserCount retrieves the total number of users in the database.
func (repository *UserRepository) GetUserCount() int {
	if repository.db == nil {
		return -1
	}

	var userCount int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
	if err != nil {
		return -1
	}
	return userCount
}

// Delete removes a user from the database by their ID.
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
