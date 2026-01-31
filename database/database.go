package database

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	cfg, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}

	// WAJIB untuk Supabase pooler
	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := stdlib.OpenDB(*cfg) // ← ini *sql.DB

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
