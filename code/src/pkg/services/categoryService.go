package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"

	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func (service *CategoryService) FindById(id uuid.UUID) *models.Category {

	res, _ := service.repo.FindById(&id)
	/* if err != nil {
		log.Println("Error finding all categories:", err)
	} */
	return res
}
func (service *CategoryService) FindAll() *[]*models.Category {

	res, _ := service.repo.FindAll()
	/* if err != nil {
		log.Println("Error finding all categories:", err)
	} */
	return res
}

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

func (service *CategoryService) Create(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Create(category); err != nil {
		return false
	}
	return true
}

func (service *CategoryService) Delete(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Delete(category); err != nil {
		return false
	}
	return true
}
func (service *CategoryService) Update(category *models.Category) bool {
	if category == nil {
		return false
	}
	if err := service.repo.Update(category); err != nil {
		return false
	}
	return true
}
