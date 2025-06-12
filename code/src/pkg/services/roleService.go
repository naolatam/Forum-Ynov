package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"os"

	"github.com/google/uuid"
)

type RoleService struct {
	repo *repositories.RoleRepository
}

func (s *RoleService) GetHighestPermRole() *models.Role {

	role, err := s.repo.FindHighestPermRole()
	if err != nil {
		return nil
	}
	return role
}

func (s *RoleService) GetMidPermRole() *models.Role {

	role, err := s.repo.FindMidPermRole()
	if err != nil {
		return nil
	}
	return role
}

func (s *RoleService) GetDefaultRole() *models.Role {

	roleStringId := os.Getenv("DEFAULT_ROLE_ID")
	if roleStringId == "" {
		return nil
	}
	roleId, err := uuid.Parse(roleStringId)
	if err != nil {
		return nil
	}

	role, err := s.repo.FindById(roleId)
	if err != nil {
		return nil
	}
	return role
}
