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

// FindById retrieves a category by its ID.
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

// FindByName retrieves a category by its name.
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

// FindByPostId retrieves all categories associated with a specific post ID.
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

// FindAll retrieves all categories from the database.
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

// DeleteCategoryByPostId deletes the association of a category with a specific post ID.
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

// AssociateCategoryToAPost associates a category with a specific post ID.
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

// Create inserts a new category into the database.
func (repository *CategoryRepository) Create(category *models.Category) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	if category == nil {
		return errors.New("category cannot be nil")
	}
	if category.ID == uuid.Nil {
		category.ID = uuid.New()
	}
	_, err := repository.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a category from the database.
func (repository *CategoryRepository) Delete(category *models.Category) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	if category == nil || category.ID == uuid.Nil {
		return errors.New("category cannot be nil or have a nil ID")
	}
	_, err := repository.db.Exec("DELETE FROM categories WHERE id = ?", category.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update modifies an existing category in the database.
func (repository *CategoryRepository) Update(category *models.Category) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	if category == nil || category.ID == uuid.Nil {
		return errors.New("category cannot be nil or have a nil ID")
	}
	_, err := repository.db.Exec("UPDATE categories SET name = ? WHERE id = ?", category.Name, category.ID)
	if err != nil {
		return err
	}
	return nil
}
