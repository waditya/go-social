package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Log the request
	log.Printf("internal server error: %s path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusInternalServerError, "The Server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the request
	log.Printf("bad request: %s path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the request
	log.Printf("Not found: %s path: %s Error: %s", r.Method, r.URL.Path, err)
	writeJSONError(w, http.StatusNotFound, err.Error())
}
