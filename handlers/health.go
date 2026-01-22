package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := HealthResponse{
		Status:  "OK",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Println("Health check responded with status OK")

}
