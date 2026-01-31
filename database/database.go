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

	// 🔑 WAJIB untuk Supabase pooler (6543)
	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := stdlib.OpenDB(*cfg)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10) // Supabase recommended kecil
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully (pgx)")
	return db, nil
}
