package repositories

import "database/sql"

func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()
	if err != nil {
		return false
	}
	return true
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &UserRepository{db: db}
}
func NewRoleRepository(db *sql.DB) *RoleRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &RoleRepository{db: db}
}
func NewSessionRepository(db *sql.DB) *SessionRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &SessionRepository{db: db}
}
