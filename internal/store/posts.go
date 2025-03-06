package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/wadiya/go-social/internal/store"
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

func (s *PostStore) GetById(ctx context.Context, postID int64) (Post, error) {
	post := &store.Post{
		ID: postID,
	}
	query := `
	SELECT * FROM posts WHERE ID=?
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.ID,
	).Scan(
		&post.Content,
		&post.Title,
		&post.UserID,
		&post.Tags,
		&post.CreatedAt,
		&post.UpdateAt,
	)

	if err != nil {
		return post, err
	}
	return post, nil
}
