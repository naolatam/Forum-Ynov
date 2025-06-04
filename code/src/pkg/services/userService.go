package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"os"

	"github.com/google/uuid"
)

type UserService struct {
	repo        *repositories.UserRepository
	sessionRepo *repositories.SessionRepository
}

func (service *UserService) FindById(id uuid.UUID) *models.User {
	user, err := service.repo.FindById(id)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByEmail(email string) *models.User {
	user, err := service.repo.FindByEmail(&email)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByUsername(username string) *models.User {
	user, err := service.repo.FindByUsername(&username)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByUsernameOrEmail(pseudo *string, email *string) *models.User {
	if pseudo == nil && email == nil {
		return nil
	}
	user, err := service.repo.FindByUsernameOrEmail(pseudo, email)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByIncompleteEmail(email string) *models.User {
	if email == "" {
		return nil
	}
	query := "%" + email + "%"
	user, err := service.repo.FindByEmail(&query)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) FindByIncompleteUsername(username string) *models.User {
	if username == "" {
		return nil
	}
	query := "%" + username + "%"
	user, err := service.repo.FindByUsername(&query)
	if err != nil {
		return nil
	}
	return user
}

func (service *UserService) GetSession(user *models.User) *models.Session {
	if user == nil || user.ID == uuid.Nil {
		return nil
	}
	session, err := service.sessionRepo.FindByUserID(user.ID)
	if err != nil {
		return nil
	}
	user.Session = *session
	return session
}

func (service *UserService) IsEmailAlreadyUse(email string) bool {
	if email == "" {
		return false
	}
	user, err := service.repo.FindByEmail(&email)
	if err != nil || user == nil {
		return false
	}
	return true
}
func (service *UserService) IsUsernameAlreadyUse(username string) bool {
	if username == "" {
		return false
	}
	user, err := service.repo.FindByUsername(&username)
	if err != nil || user == nil {
		return false
	}
	return true
}

func (service *UserService) Create(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, nil
	}
	user.ID = uuid.New()
	if user.Role_ID == uuid.Nil {
		if os.Getenv("DEFAULT_ROLE_ID") == "" {
			return nil, nil // Default role ID is not set in environment variables
		}
		// Set the default role ID from environment variable
		roleIdString := os.Getenv("DEFAULT_ROLE_ID")
		roleId, err := uuid.Parse(roleIdString)
		if err != nil {
			return nil, err // Invalid UUID format in environment variable
		}
		user.Role_ID = roleId
	}

	err := service.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
