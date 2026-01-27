package database

import (
	"context"
	"database/sql"
	"kasir-api/internal/config"
	"kasir-api/internal/utils"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgres(config *config.DatabaseConfig) (*sql.DB, func() error, error) {
	if config.URL == "" {
		log.Println("database url is empty")
		return nil, nil, utils.ErrEmptyDatabaseURL
	}

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		log.Println("failed to open database connection", err)
		return nil, nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Println("failed to ping database", err)
		db.Close()
		return nil, nil, err
	}

	closeDB := func() error {
		log.Println("closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
		return nil
	}

	return db, closeDB, nil
}
