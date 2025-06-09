package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func (repository *PostRepository) FindById(id *uuid.UUID) (*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		return &post, nil
	}
	return nil, errors.New("post not found")
}

func (repository *PostRepository) FindByTitle(title *string) (*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT * FROM posts WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		return &post, nil
	}
	return nil, errors.New("post not found")
}

func (repository *PostRepository) FindMultipleByText(text *string) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	search := "%" + *text + "%"
	rows, err := repository.db.Query(`SELECT id, title,
			CASE 
    		WHEN CHAR_LENGTH(content) > 75 THEN CONCAT(LEFT(content, 75), '...')
    		ELSE content
  			END AS content,
			picture, validated, createdAt, user_ID
			FROM posts WHERE title LIKE ? OR content LIKE ?`, search, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Post{}
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Picture, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

func (repository *PostRepository) FindByCategoryId(categoryId *uuid.UUID, limit *int) (*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 10
	if limit != nil {
		effectiveLimit = *limit
	}

	rows, err := repository.db.Query("SELECT p.* FROM posts p INNER JOIN posts_category c ON p.id = c.post_id WHERE c.category_id = ? LIMIT ?", categoryId, effectiveLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		return &post, nil
	}
	return nil, errors.New("post not found")
}

func (repository *PostRepository) FindMultipleByCategoryId(categoryId *uuid.UUID, limit *int) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 10
	if limit != nil {
		effectiveLimit = *limit
	}

	rows, err := repository.db.Query(`SELECT p.id, p.title, p.picture,
			CASE 
    		WHEN CHAR_LENGTH(p.content) > 75 THEN CONCAT(LEFT(p.content, 75), '...')
    		ELSE content
  			END AS content
		FROM posts p 
		INNER JOIN posts_category c ON p.id = c.post_id 
		WHERE c.category_id = ? AND p.validated = 1 
		LIMIT ?`, categoryId, effectiveLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var post models.Post
	var res = []*models.Post{}
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.Title, &post.Picture, &post.Content)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

func (repository *PostRepository) FindMultipleByTextAndCategory(text *string, categoryId *uuid.UUID, limit *int) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 32
	if limit != nil {
		effectiveLimit = *limit
	}
	search := "%" + *text + "%"
	rows, err := repository.db.Query(
		`SELECT p.id, p.title, p.picture,
			CASE 
    		WHEN CHAR_LENGTH(p.content) > 75 THEN CONCAT(LEFT(p.content, 75), '...')
    		ELSE content
  			END AS content
		FROM posts p 
		INNER JOIN posts_category c ON p.id = c.post_id 
		WHERE c.category_id = ? AND p.validated = 1 AND 
			(p.title LIKE ? OR p.content LIKE ?)  
		LIMIT ?`, categoryId, search, search, effectiveLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Post{}
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Picture, &post.Content)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

func (repository *PostRepository) FindLastPosts(limit *int) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 32
	if limit != nil {
		effectiveLimit = *limit
	}

	rows, err := repository.db.Query(
		`SELECT id, title, picture,
			CASE 
    		WHEN CHAR_LENGTH(content) > 75 THEN CONCAT(LEFT(content, 75), '...')
    		ELSE content
  			END AS content
		FROM posts
		WHERE validated = 1
		ORDER BY createdAt DESC
		LIMIT ?`, effectiveLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res = []*models.Post{}
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Picture, &post.Content)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

func (repository *PostRepository) GetPostCount(post *models.Post) int {
	if repository.db == nil {
		return -1
	}
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return -1
	}
	return count
}
