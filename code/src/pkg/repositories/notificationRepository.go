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
	ur *UserRepository
	pr *PostRepository
}

// FindByUserId retrieves all notifications for a specific user by their user ID.
func (repository *NotificationRepository) FindByUserId(userId uuid.UUID) (notifs *[]*models.Notification, err error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	rows, err := repository.db.Query(`SELECT id, title, description, createdAt, from_user_id, post_id FROM notifications WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifList []*models.Notification
	for rows.Next() {
		var notif models.Notification
		notif.User_ID = userId // Set the user ID for each activity
		if err := rows.Scan(&notif.ID, &notif.Title, &notif.Description, &notif.CreatedAt, &notif.FromUser_ID, &notif.Post_ID); err != nil {
			return nil, err
		}
		notif.TimeAgo = utils.TimeAgo(notif.CreatedAt) // Convert the createdAt time to a human-readable format

		if user, err := repository.ur.FindById(notif.FromUser_ID); user != nil && err == nil {
			notif.FromUser = *user // Set the FromUser field if the user is found
		} else {
			repository.Delete(notif.ID) // If the user is not found, delete the notification
			continue
		}
		if post, err := repository.pr.FindById(notif.Post_ID); post != nil && err == nil {
			notif.Post = *post // Set the Post field if the post is found
		} else {
			repository.Delete(notif.ID) // If the post is not found, delete the notification
			continue
		}
		notifList = append(notifList, &notif)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &notifList, nil
}

// Create inserts a new notification into the database.
func (repository *NotificationRepository) Create(notif *models.Notification) (success bool, err error) {

	if repository.db == nil {
		return false, errors.New("connection to database isn't established")
	}

	// Prepare the SQL statement
	_, err = repository.db.Exec(`INSERT INTO notifications(title, description, createdAt, from_user_id, user_id, post_id) 
	VALUES (?,?,?,?,?,?)`,
		notif.Title,
		notif.Description,
		notif.CreatedAt,
		notif.FromUser_ID,
		notif.User_ID,
		notif.Post_ID,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Delete removes a notification from the database by its ID.
func (repository *NotificationRepository) Delete(id int64) (success bool, err error) {
	if repository.db == nil {
		return false, errors.New("connection to database isn't established")
	}
	// Prepare the SQL statement
	_, err = repository.db.Exec(`DELETE FROM notifications WHERE id = ?`, id)
	if err != nil {
		return false, err
	}
	return true, nil

}
