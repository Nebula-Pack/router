package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// KeyHandler handles requests for /api/{key}
func KeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	if strings.Contains(key, ":") {
		// Handle key with value
		handleMetadataVersion(w, r, key)
	} else {
		// Handle just the key
		handleMetadata(w, r, key)
	}
}

func handleMetadata(w http.ResponseWriter, r *http.Request, key string) {
	// Example GET request to /metadata
	resp, err := http.Get("http://localhost:8080/metadata?key=" + key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get metadata", resp.StatusCode)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func handleMetadataVersion(w http.ResponseWriter, r *http.Request, key string) {
	// Example GET request to /metadata-version
	parts := strings.Split(key, ":")
	resp, err := http.Get("http://localhost:8080/metadata-version?key=" + parts[0] + "&version=" + parts[1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to get metadata version", resp.StatusCode)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
