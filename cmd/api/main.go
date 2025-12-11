package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	port    int
	env     string
	version string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.version, "version", "1.0.0", "API version")
	flag.Parse()

	address := fmt.Sprintf(":%d", config.port)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		config: config,
		logger: logger,
	}

	logger.Info("Starting server", "addr", address, "env", config.env)

	err := http.ListenAndServe(address, app.routes())
	if err != nil {
		logger.Error("server error", "error", err)
	}
}
