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
			return nil, server.ErrUserNotFound
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

func (s *Storage) AddLoginPass(ctx context.Context, owner string, data server.LoginPassData) (int32, error) {
	q := `INSERT INTO login_pass (owner, title, encrypted_login, encrypted_password) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int32
	err := s.db.QueryRowContext(ctx, q, owner, data.Title, data.Login, data.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to exec query: %w", err)
	}

	return id, nil
}

func (s *Storage) LoginPassList(ctx context.Context, owner string) ([]server.LoginPassData, error) {
	q := `SELECT id, title FROM login_pass WHERE owner = $1`

	rows, err := s.db.QueryContext(ctx, q, owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, server.ErrNoData
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	data := make([]server.LoginPassData, 0)
	for rows.Next() {
		d := server.LoginPassData{}
		if err := rows.Scan(&d.ID, &d.Title); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		data = append(data, d)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("failed to close rows: %w", err)
	}

	return data, nil
}

func (s *Storage) LoginPassByID(ctx context.Context, owner string, ID int32) (*server.LoginPassData, error) {
	q := `SELECT id, title, encrypted_login, encrypted_password FROM login_pass WHERE id = $1 AND owner = $2`

	data := server.LoginPassData{}
	err := s.db.QueryRowContext(ctx, q, ID, owner).Scan(&data.ID, &data.Title, &data.Login, &data.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, server.ErrNoData
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	return &data, nil
}

func (s *Storage) AddText(ctx context.Context, owner string, data server.DataText) (int32, error) {
	q := `INSERT INTO text (owner, title, encrypted_text) VALUES ($1, $2, $3) RETURNING id`

	var id int32
	err := s.db.QueryRowContext(ctx, q, owner, data.Title, data.Text).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to exec query: %w", err)
	}

	return id, nil
}

func (s *Storage) TextList(ctx context.Context, owner string) ([]server.DataText, error) {
	q := `SELECT id, title FROM text WHERE owner = $1`

	rows, err := s.db.QueryContext(ctx, q, owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, server.ErrNoData
		}

		return nil, fmt.Errorf("failed to exec query: %w", err)
	}

	data := make([]server.DataText, 0)
	for rows.Next() {
		d := server.DataText{}
		if err := rows.Scan(&d.ID, &d.Title); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		data = append(data, d)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("failed to close rows: %w", err)
	}

	return data, nil
}
