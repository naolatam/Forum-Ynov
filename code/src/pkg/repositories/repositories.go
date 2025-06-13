package repositories

import "database/sql"

// checkDBConnection checks if the database connection is valid.
func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()
	return err == nil
}

// NewUserRepository creates a new UserRepository if the database connection is valid.
func NewUserRepository(db *sql.DB) *UserRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &UserRepository{db: db}
}

// NewRecentActivityRepository creates a new RecentActivityRepository if the database connection is valid.
func NewRecentActivityRepository(db *sql.DB) *RecentActivityRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &RecentActivityRepository{db: db}
}

// NewReportRepository creates a new ReportRepository if the database connection is valid.
func NewReportRepository(db *sql.DB) *ReportRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReportRepository{db: db}
}

// NewRoleRepository creates a new RoleRepository if the database connection is valid.
func NewRoleRepository(db *sql.DB) *RoleRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &RoleRepository{db: db}
}

// NewSessionRepository creates a new SessionRepository if the database connection is valid.
func NewSessionRepository(db *sql.DB) *SessionRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &SessionRepository{db: db}
}

// NewNotificationRepository creates a new NotificationRepository if the database connection is valid.
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &NotificationRepository{db: db,
		ur: NewUserRepository(db),
		pr: NewPostRepository(db),
	}
}

// NewPostRepository creates a new PostRepository if the database connection is valid.
func NewPostRepository(db *sql.DB) *PostRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &PostRepository{db: db}
}

// NewCommentRepository creates a new CommentRepository if the database connection is valid.
func NewCommentRepository(db *sql.DB) *CommentRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &CommentRepository{db: db}
}

// NewCategoryRepository creates a new CategoryRepository if the database connection is valid.
func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &CategoryRepository{db: db}
}

// NewReactionRepository creates a new ReactionRepository if the database connection is valid.
func NewReactionRepository(db *sql.DB) *ReactionRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReactionRepository{db: db}
}
