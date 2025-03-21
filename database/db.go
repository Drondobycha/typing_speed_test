package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"typing-speed-test/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	ErrCodeUniqueViolation = "23505"
)

type Store struct {
	pool *pgxpool.Pool
}

func Newstore(connStr string) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	_, err = pool.Exec(context.Background(), `
			CREATE TABLE IF NOT EXISTS users (
					ID serial primary key,
					Username TEXT UNIQUE,
					Password TEXT,
					Email TEXT
			);
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &Store{pool: pool}, nil
}

func (s *Store) GetPasswordByUsername(username string) (string, error) {
	var password string
	err := s.pool.QueryRow(context.Background(), `
			SELECT Password FROM user WHERE Username = $1;
	`, username).Scan(&password)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", fmt.Errorf("no matches found: %w", err)
		}
		log.Printf("Database query error: %v", err)
		return "", fmt.Errorf("failed to query database: %w", err)
	}

	return password, nil
}

func (s *Store) SaveUser(user models.User) error {
	ctx := context.Background()
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Printf("failed to rollback transaction: %v", err)
		}
	}()

	_, err = tx.Exec(ctx, `
		INSERT INTO users (Username, Password, Email) VALUES ($1, $2, $3)
	`, user.Username, user.Password, user.Email)
	if err != nil {
		var DataBaseErr *pgconn.PgError
		if errors.As(err, &DataBaseErr) && DataBaseErr.Code == ErrCodeUniqueViolation {
			return fmt.Errorf("a user with this name already exists: %w", err)
		}
		return fmt.Errorf("не удаётся создать пользователя: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("transaction commit error: %v", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
