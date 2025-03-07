package store

import (
	"context"
	"database/sql"
	"errors"
)

// The Storage Strut represents tables
// We create separate interface for each table - Posts and Users

var (
	ErrNotFound = errors.New("Result not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
	}

	Users interface {
		Create(context.Context, *User) error
	}

	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}
}

// Method in Storage Package which returns Storage Structure
// or the tables
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
