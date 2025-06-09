package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"Forum-back/pkg/utils"
	"fmt"
	"html/template"
	"log"

	"github.com/google/uuid"
)

type PostService struct {
	repo *repositories.PostRepository
}

func (s *PostService) FindPostByQueryAndCategory(searchTerm string, categoryID *uuid.UUID) (*[]*models.Post, error) {
	var res *[]*models.Post
	var err error
	if searchTerm == "" && (categoryID == nil || *categoryID == uuid.Nil) { // no search term and no category
		res, err = s.repo.FindLastPosts(nil)
	}
	if searchTerm != "" && (categoryID == nil || *categoryID == uuid.Nil) { // search term only
		res, err = s.repo.FindMultipleByText(&searchTerm)
	}
	if searchTerm == "" && categoryID != nil && *categoryID != uuid.Nil { // category only
		res, err = s.repo.FindMultipleByCategoryId(categoryID, nil)
	}
	if searchTerm != "" && categoryID != nil && *categoryID != uuid.Nil { // search term and category
		res, err = s.repo.FindMultipleByTextAndCategory(&searchTerm, categoryID, nil)
	}
	if err != nil {
		log.Println("Error finding posts by text and category:", err)
	}
	if res != nil {
		for _, post := range *res {
			if post.Picture != nil {
				post.PictureBase64 = template.URL(utils.ConvertBytesToBase64(post.Picture, "image/png"))
			}
		}
	}
	return res, err
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

/* func (service *PostService) FindByTitle(title string) (*[]*models.Post, error) {
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
} */

func (service *PostService) GetPostCount() (int, error) {
	count := service.repo.GetPostCount()
	if count == -1 {
		return count, fmt.Errorf("failed to count active sessions")
	}
	return count, nil
}
