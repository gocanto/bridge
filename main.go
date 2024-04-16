package main

import (
	"github.com/gocanto/bridge/app"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", bridge.Handler)

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
