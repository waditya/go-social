package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"

	"github.com/wadiya/go-social/docs" // This is required to geb
	"github.com/wadiya/go-social/internal/store"
)

// Define structure for application
// It should include configuration and storage

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

// Define structure for capplication onfig
// It should include address, database configuration (defined separately) and environment
type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
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

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)

		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.getPostHandler)

				r.Delete("/", app.deletePostHandler)

				r.Patch("/", app.updatePostHandler)
			})
		})

		// User Route

		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.usersContextMiddleware)
				r.Get("/", app.getUserHandler)

				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})

		})
	})
	// Return the router http handler
	return r
}
func (app *application) run(mux http.Handler) error {

	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	// Create a http server using the http Server strucure
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// log.Printf("server has started at %s", app.config.addr)
	app.logger.Infow("Server has started", "addr", app.config.addr, "env", app.config.env)
	// Start the server to listen on the port and serve the traffic
	return srv.ListenAndServe()
}
