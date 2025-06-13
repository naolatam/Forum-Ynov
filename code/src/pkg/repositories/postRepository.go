package repositories

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/utils"
	"database/sql"
	"errors"
	"html/template"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

// FindById retrieves a post by its ID.
func (repository *PostRepository) FindById(id uint32) (*models.Post, error) {
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
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.Picture, &post.Validated, &post.CreatedAt, &post.User_ID)
		if err != nil {
			return nil, err
		}
		post.PictureBase64 = template.URL(utils.ConvertBytesToBase64(post.Picture, "image/png"))
		post.TimeAgo = utils.TimeAgo(post.CreatedAt)
		return &post, nil
	}
	return nil, errors.New("post not found")
}

// FindByTitle retrieves a post by its title.
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

// FindMultipleByText retrieves multiple posts that match a given text in their title or content.
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

// FindAll retrieves all posts from the database.
func (repository *PostRepository) FindAll() (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query("SELECT id, title, content, picture, validated, createdAt, user_ID FROM posts")
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
		post.PictureBase64 = template.URL(utils.ConvertBytesToBase64(post.Picture, "image/png"))
		res = append(res, &post)
	}
	return &res, nil
}

// FindWaintings retrieves all posts that are not validated (waiting for validation).
func (repository *PostRepository) FindWaintings() (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err := repository.db.Query(`
		SELECT id, title, 
			CASE 
    		WHEN CHAR_LENGTH(content) > 75 THEN CONCAT(LEFT(content, 75), '...')
    		ELSE content
  			END AS content, picture, validated, createdAt, user_ID FROM posts 
		WHERE validated = 0`)
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
		post.PictureBase64 = template.URL(utils.ConvertBytesToBase64(post.Picture, "image/png"))
		post.TimeAgo = utils.TimeAgo(post.CreatedAt)
		res = append(res, &post)
	}
	return &res, nil
}

// FindByCategoryId retrieves a single post by its category ID, with an optional limit on the number of posts.
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

// FindMultipleByCategoryId retrieves multiple posts by their category ID, with an optional limit on the number of posts.
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

// FindMultipleByTextAndCategory retrieves multiple posts that match a given text in their title or content, filtered by category ID, with an optional limit on the number of posts.
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

// FindLastPosts retrieves the last posts, with an optional limit on the number of posts.
func (repository *PostRepository) FindLastPosts(limit *int) (*[]*models.Post, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	var effectiveLimit int = 32
	if limit != nil {
		effectiveLimit = *limit
	}

	rows, err := repository.db.Query(
		`SELECT id, title, picture, content, user_id, createdAt
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
		err = rows.Scan(&post.ID, &post.Title, &post.Picture, &post.Content, &post.User_ID, &post.CreatedAt)
		post.TimeAgo = utils.TimeAgo(post.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &post)
	}
	return &res, nil
}

// GetPostCount retrieves the total number of posts in the database.
func (repository *PostRepository) GetPostCount() int {
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

// GetUserPostCount retrieves the count of posts made by a specific user.
func (repository *PostRepository) GetUserPostCount(userId *uuid.UUID) (int, error) {
	if repository.db == nil {
		return -1, errors.New("unable to connect to database")
	}
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM posts WHERE user_id = ?", userId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// UpdatePost updates an existing post in the database.
func (repository *PostRepository) UpdatePost(post *models.Post) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("UPDATE posts SET title = ?, content = ?, picture = ?, validated = ? WHERE id = ?",
		post.Title, post.Content, post.Picture, post.Validated, post.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a post from the database.
func (repository *PostRepository) Delete(post *models.Post) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM posts WHERE id = ?", post.ID)
	if err != nil {
		return err
	}
	return nil
}

// Create inserts a new post into the database, automatically assigning an ID based on the last post's ID.
func (repository *PostRepository) Create(post *models.Post) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	id, err := repository.retrieveLastId()
	if err != nil {
		return err

	}
	if id == nil {
		zeroInt := 0
		id = &zeroInt // If no ID found, start from 0
	}

	post.ID = uint32(*id + 1)

	_, err = repository.db.Exec("INSERT INTO posts (id, title, content, picture, validated, createdAt, user_ID) VALUES (?, ?, ?, ?, ?, ?, ?)",
		post.ID, post.Title, post.Content, post.Picture, post.Validated, post.CreatedAt, post.User_ID)
	if err != nil {
		return err
	}
	return nil
}

// retrieveLastId retrieves the last post ID from the database.
func (repository *PostRepository) retrieveLastId() (lastId *int, err error) {
	if repository.db == nil {
		return lastId, errors.New("connection to database isn't established")
	}
	err = repository.db.QueryRow("SELECT MAX(id) FROM posts").Scan(&lastId)
	if err != nil {
		return lastId, err
	}
	return
}
