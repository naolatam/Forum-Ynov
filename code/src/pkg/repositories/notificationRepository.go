package repositories

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/utils"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type NotificationRepository struct {
	db *sql.DB
}

func (repository *NotificationRepository) FindByUserId(userId uuid.UUID) (notifs *[]*models.Notification, err error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	rows, err := repository.db.Query(`SELECT id, title, description, createdAt FROM notifications WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifList []*models.Notification
	for rows.Next() {
		var notif models.Notification
		notif.User_ID = userId // Set the user ID for each activity
		if err := rows.Scan(&notif.ID, &notif.Title, &notif.Description, &notif.CreatedAt); err != nil {
			return nil, err
		}
		notif.TimeAgo = utils.TimeAgo(notif.CreatedAt) // Convert the createdAt time to a human-readable format
		notifList = append(notifList, &notif)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &notifList, nil
}

func (repository *NotificationRepository) Create(notif *models.Notification) (success bool, err error) {

	if repository.db == nil {
		return false, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	_, err = repository.db.Exec(`INSERT INTO notifications(title, description, createdAt, user_id) 
	VALUES (?,?,?,?)`,
		notif.Title,
		notif.Description,
		notif.CreatedAt,
		notif.User_ID,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
