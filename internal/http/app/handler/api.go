package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	appName := os.Getenv("APP_NAME")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Success",
		"app_name": appName,
	})
}
