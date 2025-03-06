package store

import (
	"context"
	"database/sql"
)

// The Storage Strut represents tables
// We create separate interface for each table - Posts and Users

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (Post, error)
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

// Method in Storage Package which returns Storage Structure
// or the tables
func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UsersStore{db},
	}
}
