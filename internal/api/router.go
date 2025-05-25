package api

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Damandeep1313/platform-backend/internal/config"
)

// NewRouter initializes and returns a new mux.Router
func NewRouter(cfg *config.Config) *mux.Router {
    router := mux.NewRouter()

    // Example middleware — you can add authentication, logging, CORS here
    // router.Use(loggingMiddleware)
    // router.Use(authMiddleware)

    // Define your API routes here

    router.HandleFunc("/health", healthCheckHandler).Methods("GET")

    // TODO: Add other routes like user, song, escrow, payments etc.
    // e.g., router.HandleFunc("/users", usersHandler).Methods("GET", "POST")

    return router
}

// healthCheckHandler is a simple endpoint to check service health
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"ok"}`))
}
