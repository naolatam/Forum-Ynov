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
