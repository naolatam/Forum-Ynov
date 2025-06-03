package repositories

import (
	"Forum-back/internal/config"
	"Forum-back/pkg/models"
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
)

type RoleRepository struct {
	db *sql.DB
}

func (repository *RoleRepository) Init() bool {
	var err error
	repository.db, err = config.OpenDBConnection()
	statusInit := true
	if err != nil {
		log.Println("[RoleRepository] Failed to connect to database")
		statusInit = false
	}
	return statusInit
}

func (repository *RoleRepository) Close() {
	err := repository.db.Close()
	if err != nil {
		log.Println("[RoleRepository] Failed to close the connection to database")
	}
}

func (repository *RoleRepository) FindByIdOrNameOrPermission(id uuid.UUID, name string, permission int) (*models.Role, error) {
if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM users WHERE id = ? OR name = ? OR permission = ?", id, name, permission)
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