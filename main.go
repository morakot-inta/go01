package main

import (
	"log"
	"morakot-inta/hello/handlers"
	"morakot-inta/hello/redis"
	"net/http"
)

func main() {

	// connect to redis
	client, err := redis.InitializedRedis("localhost:6379")
	if err != nil {
		log.Fatalf("Could not connect to Redis: %s\n", err.Error())
	}
	defer client.Close()
	log.Println("Connected to Redis successfully")

	// setup http handlers
	handlers.SetRedisClient(client)

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/auth", handlers.AuthHandler)
	http.HandleFunc("/post", handlers.PostHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
