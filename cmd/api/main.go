package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/Damandeep1313/platform-backend/internal/config"    // **note: no 'internal' here**
    "github.com/Damandeep1313/platform-backend/internal/api"
)


func main() {
    cfg, err := config.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    router := api.NewRouter(cfg)

    addr := fmt.Sprintf(":%d", cfg.Server.Port)
    log.Printf("Starting server on %s", addr)

    if err := http.ListenAndServe(addr, router); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
