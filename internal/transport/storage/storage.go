package storage

import (
	"context"
	"database/sql"

	"github.com/lks-go/pass-keeper/internal/service"
)

// New is Storage constructor
func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// Storage is storage main struct
type Storage struct {
	db *sql.DB
}

func (s *Storage) Register(ctx context.Context, u service.User) (string, error) {
	//TODO implement me
	panic("implement me")
}
