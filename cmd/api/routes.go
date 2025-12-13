package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)
	mux.HandleFunc("GET /v1/users/{id}", app.showUserHandler)
	mux.HandleFunc("POST /v1/users", app.createUserHandler)
	mux.HandleFunc("PATCH /v1/users/{id}", app.updateUserHandler)
	mux.HandleFunc("DELETE /v1/users/{id}", app.deleteUserHandler)

	return mux
}
