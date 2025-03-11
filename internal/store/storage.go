package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// The Storage Strut represents tables
// We create separate interface for each table - Posts and Users

var (
	ErrNotFound          = errors.New("Result not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		DeleteByID(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}

	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}

	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
		DeleteByPostID(context.Context, int64) error
		Create(context.Context, *Comment) error
	}
	Followers interface {
		Follow(ctx context.Context, followerId int64, userID int64) error
		Unfollow(ctx context.Context, followerId int64, userID int64) error
	}
}

// Method in Storage Package which returns Storage Structure
// or the tables
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
