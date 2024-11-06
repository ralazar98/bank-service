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

type TechRoute struct {
	version string
}

func NewTechRouteHandler() *TechRoute {
	return &TechRoute{
		version: os.Getenv("VERSION"),
	}
}

func (t *TechRoute) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	health := HealthResponse{
		Status: "UP",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func (t *TechRoute) VersionHandler(w http.ResponseWriter, r *http.Request) {

	version := VersionResponse{
		Version: t.version,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}
