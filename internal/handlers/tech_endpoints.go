package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status string `json:"status"`
}
type VersionResponse struct {
	Version string `json:"version"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	health := HealthResponse{
		Status: "UP",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {

	version := VersionResponse{
		Version: os.Getenv("APP_VERSION"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}
