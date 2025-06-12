package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type CategoryRepository struct {
	db *sql.DB
}

func (repository *CategoryRepository) FindById(id *uuid.UUID) (*models.Category, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM categories WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category models.Category
	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}
	return nil, errors.New("category not found")
}

func (repository *CategoryRepository) FindByName(name *string) (*models.Category, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM categories WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category models.Category
	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		return &category, nil
	}
	return nil, errors.New("category not found")
}

func (repository *CategoryRepository) FindByPostId(postId uint32) (*[]*models.Category, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT c.* FROM categories c INNER JOIN posts_category pc ON c.id = pc.category_id WHERE pc.post_id = ?", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Category{}
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, &category)
	}
	return &res, nil
}

func (repository *CategoryRepository) FindAll() (*[]*models.Category, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Category{}
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, &category)
	}
	return &res, nil
}

func (repository *CategoryRepository) DeleteCategoryByPostId(postId uint32) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM posts_category WHERE post_id = ?", postId)
	if err != nil {
		return err
	}
	return nil
}

func (repository *CategoryRepository) AssociateCategoryToAPost(categoryId uuid.UUID, postId uint32) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("INSERT INTO posts_category (category_id, post_id) VALUES (?, ?)", categoryId, postId)
	if err != nil {
		return err
	}

	return nil
}
