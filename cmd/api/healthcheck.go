package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     app.config.version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "The server encountered a problem", http.StatusInternalServerError)
	}
}
