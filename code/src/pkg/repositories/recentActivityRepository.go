package repositories

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/utils"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type RecentActivityRepository struct {
	db *sql.DB
}

// FindByUserId retrieves recent activities for a specific user by their user ID.
func (repository *RecentActivityRepository) FindByUserId(userId uuid.UUID) (activities *[]*models.RecentActivity, err error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	rows, err := repository.db.Query(`SELECT action, details, subTitle, post_id, createdAt FROM recent_activity WHERE user_id = ? ORDER BY createdAt DESC`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activitiesList []*models.RecentActivity
	for rows.Next() {
		var activity models.RecentActivity
		activity.User_ID = userId // Set the user ID for each activity
		if err := rows.Scan(&activity.Action, &activity.Details, &activity.SubTitle, &activity.Post_ID, &activity.CreatedAt); err != nil {
			return nil, err
		}
		activity.TimeAgo = utils.TimeAgo(activity.CreatedAt) // Convert the createdAt time to a human-readable format
		activitiesList = append(activitiesList, &activity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &activitiesList, nil
}

// Create inserts a new recent activity into the database.
func (repository *RecentActivityRepository) Create(activity *models.RecentActivity) (success bool, err error) {

	if repository.db == nil {
		return false, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	_, err = repository.db.Exec(`INSERT INTO recent_activity(action, details, subTitle, user_id, post_id, createdAt) 
	VALUES (?,?,?,?,?,?)`,
		activity.Action,
		activity.Details,
		activity.SubTitle,
		activity.User_ID,
		activity.Post_ID,
		activity.CreatedAt,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
