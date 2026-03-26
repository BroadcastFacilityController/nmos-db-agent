package database

import (
	"context"
	"database/sql"

	"github.com/BroadcastFacilityController/nmos-db-agent/database/migrations"
	"github.com/golang-migrate/migrate/v4"
	pgxv5 "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *DB

func NewConnection(ctx context.Context, user string, password string, address string, port string) (*DB, error) {
	// Connect to database
	url := "postgres://" + user + ":" + password + "@" + address + ":" + port + "/nmos"
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	// Ensure we connected successfully
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	// Do migration
	err = doMigrations(url)
	if err != nil {
		return nil, err
	}

	dbtemp := DB{
		db: conn,
	}

	db = &dbtemp
	return db, nil
}

func doMigrations(dsn string) error {
	// Get native SQL driver
	stddb, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer stddb.Close()

	err = stddb.Ping()
	if err != nil {
		return err
	}

	// Get migration driver
	cfg := pgxv5.Config{}
	migrationDriver, err := pgxv5.WithInstance(stddb, &cfg)
	if err != nil {
		return err
	}

	sourceDriver, err := iofs.New(migrations.EmbeddedFS, ".")
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", sourceDriver, "pgx/v5", migrationDriver)
	if err != nil {
		return err
	}

	// Do migration
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

type DB struct {
	db *pgxpool.Pool
}

func Close() {
	db.db.Close()
}
