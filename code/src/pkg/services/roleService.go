package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
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
