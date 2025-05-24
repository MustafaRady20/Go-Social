package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context)
	}
	Users interface {
		Create(context.Context)
	}
}

func NewPostgresStorage(db *sql.DB) error {
	return Storage{
		Posts: &PostgresPosts{
			db: db,		
	}
}
