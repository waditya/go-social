package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Map format : map[<<data type of key>>]<<data type of value>>
	// A map with key as string and value as Boolean will be
	// map[string]bool

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
