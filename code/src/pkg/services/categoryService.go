package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
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
