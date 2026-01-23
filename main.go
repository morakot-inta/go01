package main

import (
	"context"
	"log"
	"morakot-inta/hello/handlers"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func initializedRedis(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {

	// connect to redis
	client, err := initializedRedis("localhost:6379")
	if err != nil {
		log.Fatalf("Could not connect to Redis: %s\n", err.Error())
	}
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

	defer client.Close()

}
