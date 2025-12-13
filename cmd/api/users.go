package main

import (
	"net/http"
	"strconv"
)

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := app.models.Users.Get(id)
	if err != nil {
		app.logger.Error("Failed to get user", "error", err, "id", id)

		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Error("write json error", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
