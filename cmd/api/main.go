package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	"titan/internal/data"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type db struct {
	dsn string
}
type config struct {
	port    int
	env     string
	version string
	db      db
}

type application struct {
	config config
	logger *slog.Logger
	db     *sql.DB
	models data.Models
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.version, "version", "1.0.0", "API version")
	flag.StringVar(&config.db.dsn, "db-dsn", "", "PostgreSQL DSN")
	flag.Parse()

	address := fmt.Sprintf(":%d", config.port)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(config)

	if err != nil {
		logger.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	app := &application{
		config: config,
		logger: logger,
		db:     db,
		models: data.NewModels(db),
	}

	defer db.Close()

	logger.Info("Starting server", "addr", address, "env", config.env)

	err = http.ListenAndServe(address, app.routes())
	if err != nil {
		logger.Error("server error", "error", err)
	}
}

func openDB(config config) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
