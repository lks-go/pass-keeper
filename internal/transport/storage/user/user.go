package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/lks-go/pass-keeper/internal/service"
)

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db *sql.DB
}

type user struct {
	ID           string
	Login        string
	PasswordHash string
}

func (s *Storage) AddUser(ctx context.Context, login string, passwordHash string) (string, error) {
	q := `INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id`

	id := ""
	err := s.db.QueryRowContext(ctx, q, login, passwordHash).Scan(&id)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			if err.Code == pgerrcode.UniqueViolation {
				return "", service.ErrAlreadyExists
			}
		}

		return "", fmt.Errorf("failed to exec query: %w", err)
	}

	return id, nil
}

func (s *Storage) UserByLogin(ctx context.Context, login string) (*service.User, error) {
	q := `SELECT id, login, password_hash FROM users WHERE login = $1;`

	u := user{}
	if err := s.db.QueryRowContext(ctx, q, login).Scan(&u.ID, &u.Login, &u.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, service.ErrNotFound
		}

		return nil, fmt.Errorf("query row error: %w", err)
	}

	su := service.User{
		ID:           u.ID,
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
	}

	return &su, nil
}
