package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/users/{id}", app.showUserHandler)

	return mux
}
