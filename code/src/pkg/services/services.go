package services

import (
	"Forum-back/pkg/repositories"
	"database/sql"
)

func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()

	return err == nil
}

func NewCategoryService(db *sql.DB) *CategoryService {
	if !checkDBConnection(db) {
		return nil
	}
	return &CategoryService{
		repo: repositories.NewCategoryRepository(db),
	}
}

func NewCommentService(db *sql.DB) *CommentService {
	if !checkDBConnection(db) {
		return nil
	}
	return &CommentService{
		repo: repositories.NewCommentRepository(db),
		us:   repositories.NewUserRepository(db),
		ps:   repositories.NewPostRepository(db),
	}
}

func NewNotificationService(db *sql.DB) *NotificationService {
	if !checkDBConnection(db) {
		return nil
	}
	return &NotificationService{
		repo: repositories.NewNotificationRepository(db),
		ur:   repositories.NewUserRepository(db),
	}
}

func NewPostService(db *sql.DB) *PostService {
	if !checkDBConnection(db) {
		return nil
	}
	return &PostService{
		repo: repositories.NewPostRepository(db),
		ur:   repositories.NewUserRepository(db),
	}
}

func NewReactionService(db *sql.DB) *ReactionService {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReactionService{
		repo: repositories.NewReactionRepository(db),
	}
}

func NewRecentActivityService(db *sql.DB) *RecentActivityService {
	if !checkDBConnection(db) {
		return nil
	}
	return &RecentActivityService{
		repo: repositories.NewRecentActivityRepository(db),
		ur:   repositories.NewUserRepository(db),
		pr:   repositories.NewPostRepository(db),
	}
}

func NewRoleService(db *sql.DB) *RoleService {
	if !checkDBConnection(db) {
		return nil
	}
	return &RoleService{
		repo: repositories.NewRoleRepository(db),
	}
}

func NewSessionService(db *sql.DB) *SessionService {
	if !checkDBConnection(db) {
		return nil
	}
	return &SessionService{
		repo:     repositories.NewSessionRepository(db),
		userRepo: repositories.NewUserRepository(db),
	}
}

func NewUserService(db *sql.DB) *UserService {
	if !checkDBConnection(db) {
		return nil
	}
	return &UserService{
		repo:        repositories.NewUserRepository(db),
		sessionRepo: repositories.NewSessionRepository(db),
		roleRepo:    repositories.NewRoleRepository(db),
	}
}
