package repositories

import "database/sql"

func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()
	return err == nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &UserRepository{db: db}
}

func NewRecentActivityRepository(db *sql.DB) *RecentActivityRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &RecentActivityRepository{db: db}
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReportRepository{db: db}
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

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &NotificationRepository{db: db,
		ur: NewUserRepository(db),
		pr: NewPostRepository(db),
	}
}

func NewPostRepository(db *sql.DB) *PostRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &PostRepository{db: db}
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &CommentRepository{db: db}
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &CategoryRepository{db: db}
}

func NewReactionRepository(db *sql.DB) *ReactionRepository {
	if !checkDBConnection(db) {
		return nil
	}
	return &ReactionRepository{db: db}
}
