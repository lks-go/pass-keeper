package setup

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func DB(dsn string) (*sql.DB, error) {
	fmt.Printf("dsn: %s\n", dsn)
	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database after connect: %w", err)
	}

	return pool, nil
}
