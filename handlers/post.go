package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type PostResponse struct {
	Message string `json:"message"`
}

type PostRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func SetRedisClient(client *redis.Client) {
	client = client
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("PostHandler called")

	// check method is POST
	if r.Method == http.MethodPost {

		var postReq PostRequestBody

		err := json.NewDecoder(r.Body).Decode(&postReq)
		if err != nil {
			log.Printf("Error decoding post request: %s\n", err.Error())

			res := PostResponse{
				Message: "Invalid request payload",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Printf("Received post: Title=%s, Content=%s\n", postReq.Title, postReq.Content)

		// publish to redis stream
		client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})

		streamID, err := client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: "posts",
			ID:     "*",
			Values: map[string]interface{}{
				"title":     postReq.Title,
				"content":   postReq.Content,
				"timestamp": time.Now().Unix(),
			},
		}).Result()

		defer client.Close()

		if err != nil {
			log.Printf("Error adding post to Redis stream: %s\n", err.Error())

			res := PostResponse{
				Message: "Failed to create post",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Printf("Post added to Redis stream with ID: %s\n", streamID)

		res := PostResponse{
			Message: "Post created successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
		return
	}

	// if method is not POST
	res := PostResponse{
		Message: "Method Not Allowed",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(res)
}
