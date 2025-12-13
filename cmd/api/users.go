package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"titan/internal/data"
)

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid user ID"))
		return
	}

	user, err := app.models.Users.Get(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/users/%d", user.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// TODO : does not notified client that user created successfully
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid user ID"))
		return
	}

	user, err := app.models.Users.Get(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Name  *string `json:"name"`
		Email *string `json:"email"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	err = app.models.Users.Update(user)
	if err != nil {
		if err.Error() == "edit conflict" {
			app.errorResponse(w, r, http.StatusConflict, "unable to update due to edit conflict, please try again")
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid user ID"))
		return
	}

	err = app.models.Users.Delete(id)
	if err != nil {
		if err.Error() == "record not found" {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "user successfully deleted"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
