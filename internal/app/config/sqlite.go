package config

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/romik1505/youtubeThumbnails/internal/app/store"
	"log"
)

// NewSqliteConnection .
func NewSqliteConnection(ctx context.Context, connString string) store.Storage {
	con, err := sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatalln("database connection err: %w", err)
	}

	if err := con.Ping(); err != nil {
		log.Fatalln(err)
	}

	return store.Storage{
		DB: con,
	}
}
