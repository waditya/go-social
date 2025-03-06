package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/wadiya/go-social/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		// writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// userId := 1

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

}

// *Applications invokes getPostHandler

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	// Dummmy postID

	// var postID int64 = 2

	// Extract URLParam to get postID

	idParam := chi.URLParam(r, "postID")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		// writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get the context from the request

	ctx := r.Context()

	// Retrieve the post from the Posts table from the store
	// created from the app
	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
			// writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			app.internalServerError(w, r, err)
			// writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		// writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
