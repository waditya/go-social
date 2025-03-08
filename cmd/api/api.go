package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wadiya/go-social/internal/store"
)

// Define structure for application
// It should include configuration and storage

type application struct {
	config config
	store  store.Storage
}

// Define structure for capplication onfig
// It should include address, database configuration (defined separately) and environment
type config struct {
	addr string
	db   dbConfig
	env  string
}

// Define structure for database configuration
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// Define a mount() function for pointer to application struct
// It must return an http handler
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.getPostHandler)

				r.Delete("/", app.deletePostHandler)

				r.Patch("/", app.updatePostHandler)
			})
		})
	})
	// Return the router http handler
	return r
}
func (app *application) run(mux http.Handler) error {

	// Create a http server using the http Server strucure
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)
	// Start the server to listen on the port and serve the traffic
	return srv.ListenAndServe()
}
