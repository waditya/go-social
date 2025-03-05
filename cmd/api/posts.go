package main

import (
	"net/http"

	"github.com/wadiya/go-social/internal/store"
)

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// userId := 1

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
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
