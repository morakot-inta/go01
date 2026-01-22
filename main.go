package main

import (
	"log"
	"morakot-inta/hello/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/auth", handlers.AuthHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
