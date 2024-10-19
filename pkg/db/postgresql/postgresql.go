package db

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

var TemplateFs embed.FS

type Client interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewClient(ctx context.Context, maxAttempts int, dsn string) (pool *pgxpool.Pool, err error) {
	if err = doWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}

		return nil
	}, maxAttempts, 5*time.Second); err != nil {
		return nil, fmt.Errorf("failed to create connection pool after %d attempts: %w", maxAttempts, err)
	}

	return pool, nil
}

func MigrationsUp(pool *pgxpool.Pool, migrationFolder string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(pool)
	if err := goose.Up(db, migrationFolder); err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}

	return nil
}

func doWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return err
}
