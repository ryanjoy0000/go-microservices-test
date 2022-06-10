package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (config *BrokerAppConfig) routes() http.Handler {
	// Create chi mux
	chiMux := chi.NewRouter()

	// Specify what & who is allowed to connect to http server
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	corsHandler := cors.Handler(corsOptions)
	chiMux.Use(corsHandler)

	// --- Routes ---
	// Check uptime by load-balancers or other services before hitting other routes
	chiMux.Use(middleware.Heartbeat("/ping"))
	chiMux.Post("/", config.BrokerPost)

	return chiMux
}
