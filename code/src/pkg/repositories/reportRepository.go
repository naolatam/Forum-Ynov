package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"
)

type ReportRepository struct {
	db *sql.DB
}

// FindAll retrieves all reports from the database, ordered by the reportedAt timestamp.
func (repository *ReportRepository) FindAll() (*[]*models.Report, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM reports ORDER BY reportedAt ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*models.Report
	if rows.Next() {
		var report models.Report
		err = rows.Scan(&report.ID, &report.Post_id, &report.User_id, &report.ReportAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &report)
	}
	return &res, nil
}

// Create inserts a new report into the database.
func (repository *ReportRepository) Create(report *models.Report) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("INSERT INTO reports ( id, post_id, user_id, reportedAt) VALUES ( ?, ?, ?, ?)",
		report.ID, report.Post_id, report.User_id, report.ReportAt)
	return err
}

// Delete removes a report from the database by its ID.
func (repository *ReportRepository) Delete(report *models.Report) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM reports WHERE id = ?", report.ID)
	return err
}
