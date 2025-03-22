package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the authorization header
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
				return
			}
			// Parse it --> Get the base64

			parts := strings.Split(authHeader, " ")

			if len(parts) != 2 || parts[0] != "Basic" {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("authorization header is malformed"))
				return
			}

			// Decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])

			if err != nil {
				app.unauthorizedErrorResponse(w, r, err)
				return
			}

			username := app.config.auth.basic.user
			pass := app.config.auth.basic.pass

			// decoded is a slice of bytes. It consists decoded username:password

			creds := strings.SplitN(string(decoded), ":", 2)

			if len(creds) != 2 || creds[0] != username || creds[1] != pass {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			// Check the credentials

			next.ServeHTTP(w, r)
		})
	}
}
