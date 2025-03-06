package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// Specify the model for Post
type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB //Pointer to database connection
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdateAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postID int64) (*Post, error) {

	query := `
	SELECT id, title, user_id, content, created_at, updated_at, tags 
	FROM posts 
	WHERE ID=$1
	`

	var post Post

	err := s.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.UserID,
		&post.Content,
		&post.CreatedAt,
		&post.UpdateAt,
		pq.Array(post.Tags),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &post, nil
}
