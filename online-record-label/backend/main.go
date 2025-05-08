package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Enable CORS for cross-origin requests
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	http.HandleFunc("/api/records", recordsHandler)
	http.HandleFunc("/api/artists", artistsHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func recordsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "application/json")
	records := []string{"Record 1", "Record 2", "Record 3"}
	json.NewEncoder(w).Encode(records)
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "application/json")
	artists := []string{"Artist 1", "Artist 2", "Artist 3"}
	json.NewEncoder(w).Encode(artists)
}
