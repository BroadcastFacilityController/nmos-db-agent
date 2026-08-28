package nmosdb

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	Pool    *pgxpool.Pool
	Queries *Queries
}

type Config struct {
	DatabaseURL   string
	RunMigrations bool
}

func Open(ctx context.Context, cfg Config) (*DB, error) {
	if cfg.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	stdb := stdlib.OpenDBFromPool(pool)

	if cfg.RunMigrations {
		err := RunMigrations(stdb)
		if err != nil {
			return nil, err
		}
	}

	return &DB{
		Pool:    pool,
		Queries: New(pool),
	}, nil
}

func (d *DB) Close() {
	if d == nil || d.Pool == nil {
		return
	}
	d.Pool.Close()
}

func RunMigrations(db *sql.DB) error {
	// Do migrations
	goose.SetBaseFS(migrationsFS)

	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
