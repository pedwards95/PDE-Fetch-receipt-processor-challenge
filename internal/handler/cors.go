package handler

import "github.com/go-chi/cors"

// setupCors basic cors setup, doesn't currently do much
func (hl *Handler) setupCors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	return cors
}
