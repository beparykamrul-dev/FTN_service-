package main

import (
    "log"
    "net/http"

    "github.com/beparykamrul-dev/FTN_service-/internal/frontier"
)

func main() {
    registry := frontier.NewRegistry()
    registry.Register(frontier.Module{ID: "frontier-002", Name: "Physics Data Ingestion & Source Adapter Engine", Domain: "frontier-research", Classification: frontier.Experimental, Status: frontier.Active})

    mux := http.NewServeMux()
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte("ok\n")) })
    mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte("ready\n")) })
    mux.HandleFunc("/api/v1/frontier/modules", registry.Handler)

    log.Println("FTN service listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
