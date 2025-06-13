package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"log"

	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

// FindById retrieves a category by its ID.
func (service *CategoryService) FindById(id uuid.UUID) *models.Category {

	res, err := service.repo.FindById(&id)
	if err != nil {
		log.Println("Error finding all categories:", err)
	}
	return res
}

// FindAll retrieves all categories.
func (service *CategoryService) FindAll() *[]*models.Category {

	res, err := service.repo.FindAll()
	if err != nil {
		log.Println("Error finding all categories:", err)
	}
	return res
}

// FindByPostId retrieves all categories associated with a specific post.
func (s *CategoryService) FindByPostId(post *models.Post) (*[]*models.Category, error) {
	if post == nil {
		return nil, nil
	}
	c, err := s.repo.FindByPostId((post.ID))
	if err != nil {
		return nil, err
	}
	post.Categories = *c
	return c, nil
}

// Create adds a new category to the repository.
func (service *CategoryService) Create(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Create(category); err != nil {
		return false
	}
	return true
}

// Delete removes a category from the repository.
func (service *CategoryService) Delete(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Delete(category); err != nil {
		return false
	}
	return true
}

// Update modifies an existing category in the repository.
func (service *CategoryService) Update(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Update(category); err != nil {
		return false
	}
	return true
}
