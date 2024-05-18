package main

import (
	"log"
	"net/http"

	"github.com/omardakelbab/api_v2/pkg/api"
)

func main() {
	port := ":8080"
	log.Printf("Starting server on %s", port)

	// Initialize API handlers
	http.HandleFunc("/users", api.HandleUsers)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
