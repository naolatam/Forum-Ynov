package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ReportService struct {
	repository *repositories.ReportRepository
	ur         *repositories.UserRepository
	pr         *repositories.PostRepository
}

// FindAll retrieves all reports from the repository and populates user and post information.
func (s *ReportService) FindAll() (*[]*models.Report, error) {
	reports, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, report := range *reports {
		if report.User_id != uuid.Nil {
			user, err := s.ur.FindById(report.User_id)
			if err != nil {
				return nil, err
			}
			report.User = *user
		}
		post, err := s.pr.FindById(report.Post_id)
		if err != nil {
			return nil, err
		}
		post.Content = post.Content[:max(75, len(post.Content))] // Truncate content to 75 characters
		report.Post = *post

	}

	return reports, nil
}

// Create adds a new report to the repository.
func (s *ReportService) Create(report *models.Report) error {

	report.ID = uuid.New()
	report.ReportAt = time.Now()
	if report.Post_id == 0 {
		return errors.New("post_id cannot be 0")
	}
	if report.User_id == uuid.Nil {
		return errors.New("user_id cannot be nil")
	}

	if err := s.repository.Create(report); err != nil {
		return err
	}
	return nil

}

// Delete removes a report from the repository.
func (s *ReportService) Delete(report *models.Report) error {
	if report == nil {
		return errors.New("report cannot be nil")
	}
	if err := s.repository.Delete(report); err != nil {
		return err
	}
	return nil
}
