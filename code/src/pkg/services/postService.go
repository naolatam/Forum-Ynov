package services

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"Forum-back/pkg/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	repo *repositories.PostRepository
	ur   *repositories.UserRepository
	cr   *repositories.CategoryRepository
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
func (service *PostService) FindById(id uint32) (*models.Post, error) {
	post, err := service.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FetchUserId(post *models.Post) (*models.User, error) {
	if post == nil || post.User_ID == uuid.Nil {
		return nil, fmt.Errorf("post or user ID is nil")
	}
	user, err := s.ur.FindById(post.User_ID)
	if err != nil {
		return nil, err
	}
	post.User = *user
	return user, nil
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

func (service *PostService) GetUserPostCount(user *models.User) int {
	if user == nil || user.ID == uuid.Nil {
		return -1
	}
	count, _ := service.repo.GetUserPostCount(&user.ID)

	return count
}

func (service *PostService) UpdateCategoryFromList(categories []string, post *models.Post) dtos.ErrorPageDto {
	for _, categoryID := range categories {
		if categoryID == "" {
			return dtos.ErrorPageDto{
				Header:  dtos.HeaderDto{},
				Message: "Category ID cannot be empty.",
				Details: "Please provide a valid category ID.",
				Code:    http.StatusBadRequest,
			}
		}
		validCategoryId, err := uuid.Parse(categoryID)
		if err != nil {
			return dtos.ErrorPageDto{
				Header:  dtos.HeaderDto{},
				Message: "Invalid category ID format.",
				Details: fmt.Sprintf("The provided category ID '%s' is not a valid UUID.", categoryID),
				Code:    http.StatusBadRequest,
			}
		}
		c, err := service.cr.FindById(&validCategoryId)
		if c != nil && err == nil {
			post.Categories = append(post.Categories, c)
		}
	}

	if err := service.UpdateCategory(post); err != nil {
		return dtos.ErrorPageDto{
			Header:  dtos.HeaderDto{},
			Message: "Error while updating post category",
			Details: fmt.Sprintf("Error while updating post category: %s", err.Error()),
			Code:    http.StatusInternalServerError,
		}
	}
	return dtos.ErrorPageDto{
		Code: http.StatusOK,
	}
}

func (service *PostService) UpdateCategory(post *models.Post) error {
	if post == nil || post.ID == 0 {
		return fmt.Errorf("post is nil or has no ID")
	}

	if err := service.cr.DeleteCategoryByPostId(post.ID); err != nil {
		return fmt.Errorf("error updating existing categories for post ID %d: %s", post.ID, err)
	}

	for _, category := range post.Categories {
		if category == nil || category.ID == uuid.Nil {
			continue
		}
		if err := service.cr.AssociateCategoryToAPost(category.ID, post.ID); err != nil {
			return fmt.Errorf("error adding category %s to post ID %d: %s", category.Name, post.ID, err)
		}
	}

	return nil
}

func (service *PostService) UpdatePost(post *models.Post) error {
	if post == nil || post.ID == 0 {
		return fmt.Errorf("post is nil or has no ID")
	}

	if err := service.repo.UpdatePost(post); err != nil {
		return fmt.Errorf("error updating post ID %d: %s", post.ID, err)
	}

	return nil
}

func (service *PostService) Delete(post *models.Post) error {
	if post == nil || post.ID == 0 {
		return fmt.Errorf("post is nil or has no ID")
	}

	if err := service.repo.Delete(post); err != nil {
		return fmt.Errorf("error deleting post ID %d: %s", post.ID, err)
	}
	return nil
}

func (service *PostService) Create(post *models.Post) error {
	if post == nil {
		return fmt.Errorf("post is nil")
	}

	post.CreatedAt = time.Now()
	if post.Picture == nil {
		post.Picture = utils.GetDefaultAvatar()
	}

	if err := service.repo.Create(post); err != nil {
		return fmt.Errorf("error creating post: %s", err)
	}

	if len(post.Categories) > 0 {
		if err := service.UpdateCategory(post); err != nil {
			return fmt.Errorf("error associating categories with post: %s", err)
		}
	}

	return nil
}
