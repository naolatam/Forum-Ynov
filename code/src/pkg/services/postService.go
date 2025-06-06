package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"fmt"

	"github.com/google/uuid"
)

type PostService struct {
	repo *repositories.PostRepository
}

func (service *PostService) FindById(id uuid.UUID) (*models.Post, error) {
	if id == uuid.Nil {
		return nil, nil
	}
	post, err := service.repo.FindById(&id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (service *PostService) FindByTitle(title string) (*[]*models.Post, error) {
	if title == "" {
		return nil, nil
	}
	posts, err := service.repo.FindByTitle(&title)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (service *PostService) FindByCategoryId(categoryId uuid.UUID, limit *int) (*[]*models.Post, error) {
	if categoryId == uuid.Nil {
		return nil, nil
	}
	posts, err := service.repo.FindByCategoryId(&categoryId, limit)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (service *PostService) GetPostCount() (int, error) {
	count := service.repo.GetPostCount()
	if count == -1 {
		return count, fmt.Errorf("failed to count active sessions")
	}
	return count, nil
}








