package services

import (
	"Forum-back/pkg/repositories"
	"database/sql"
)

// checkDBConnection checks if the database connection is valid.
func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()

	return err == nil
}

// NewCategoryService creates a new CategoryService if the database connection is valid.
func NewCategoryService(db *sql.DB) *CategoryService {
	if !checkDBConnection(db) {
		return nil
	}
	return &CategoryService{
		repo: repositories.NewCategoryRepository(db),
	}
}

// NewCommentService creates a new CommentService if the database connection is valid.
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

// NewNotificationService creates a new NotificationService if the database connection is valid.
func NewNotificationService(db *sql.DB) *NotificationService {
	if !checkDBConnection(db) {
		return nil
	}
	return &NotificationService{
		repo: repositories.NewNotificationRepository(db),
		ur:   repositories.NewUserRepository(db),
	}
}

// NewPostService creates a new PostService if the database connection is valid.
func NewPostService(db *sql.DB) *PostService {
	if !checkDBConnection(db) {
		return nil
	}
	return &PostService{
		repo:         repositories.NewPostRepository(db),
		ur:           repositories.NewUserRepository(db),
		cr:           repositories.NewCategoryRepository(db),
		roleRepo:     repositories.NewRoleRepository(db),
		reactionRepo: repositories.NewReactionRepository(db),
	}
}

// NewReactionService creates a new ReactionService if the database connection is valid.
func NewReactionService(db *sql.DB) *ReactionService {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReactionService{
		repo: repositories.NewReactionRepository(db),
	}
}

// NewRecentActivityService creates a new RecentActivityService if the database connection is valid.
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

// NewReportService creates a new ReportService if the database connection is valid.
func NewReportService(db *sql.DB) *ReportService {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReportService{
		repository: repositories.NewReportRepository(db),
		ur:         repositories.NewUserRepository(db),
		pr:         repositories.NewPostRepository(db),
	}
}

// NewRoleService creates a new RoleService if the database connection is valid.
func NewRoleService(db *sql.DB) *RoleService {
	if !checkDBConnection(db) {
		return nil
	}
	return &RoleService{
		repo: repositories.NewRoleRepository(db),
	}
}

// NewSessionService creates a new SessionService if the database connection is valid.
func NewSessionService(db *sql.DB) *SessionService {
	if !checkDBConnection(db) {
		return nil
	}
	return &SessionService{
		repo:     repositories.NewSessionRepository(db),
		userRepo: repositories.NewUserRepository(db),
	}
}

// NewUserService creates a new UserService if the database connection is valid.
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
