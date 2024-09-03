package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/lks-go/pass-keeper/internal/service/server"
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

func (s *Storage) RegisterUser(ctx context.Context, login string, passwordHash string) (string, error) {
	q := `INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id`

	id := ""
	err := s.db.QueryRowContext(ctx, q, login, passwordHash).Scan(&id)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			if err.Code == pgerrcode.UniqueViolation {
				return "", server.ErrAlreadyExists
			}
		}

		return "", fmt.Errorf("failed to exec query: %w", err)
	}

	return id, nil
}

func (s *Storage) UserByLogin(ctx context.Context, login string) (*server.User, error) {
	q := `SELECT id, login, password_hash FROM users WHERE login = $1;`

	u := user{}
	if err := s.db.QueryRowContext(ctx, q, login).Scan(&u.ID, &u.Login, &u.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, server.ErrNotFound
		}

		return nil, fmt.Errorf("query row error: %w", err)
	}

	su := server.User{
		ID:           u.ID,
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
	}

	return &su, nil
}

func (s *Storage) AddLoginPass(ctx context.Context, owner string, data server.Data) error {
	q := `INSERT INTO login_pass (owner, title, encrypted_login, encrypted_password) VALUES ($1, $2, $3, $4)`

	_, err := s.db.ExecContext(ctx, q, owner, data.Title, data.Login, data.Password)
	if err != nil {
		return fmt.Errorf("failed to exec query: %w", err)
	}

	return nil
}

func (s *Storage) LoginPassList(ctx context.Context, owner string) ([]server.Data, error) {
	q := `SELECT id, title, encrypted_login, encrypted_password FROM login_pass WHERE owner = $1`

	rows, err := s.db.QueryContext(ctx, q, owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, server.ErrNotFound
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	data := make([]server.Data, 0)
	for rows.Next() {
		d := server.Data{}
		if err := rows.Scan(&d.ID, &d.Title, &d.Login, &d.Password); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		data = append(data, d)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("failed to close rows: %w", err)
	}

	return data, nil
}
