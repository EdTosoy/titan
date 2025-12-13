package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Error("error", "error", err, "method", r.Method, "uri", r.URL.RequestURI())

}
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	data := envelope{"error": message}
	err := app.writeJSON(w, status, data, nil)

	if err != nil {
		app.logError(r, err)

		w.WriteHeader((http.StatusInternalServerError))
	}
}
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
