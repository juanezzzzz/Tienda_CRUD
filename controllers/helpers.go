package controllers

import (
	"encoding/json"
	"net/http"
)

// Helper respuesta JSON
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
