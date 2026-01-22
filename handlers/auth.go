package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("AuthHandler called")

	// check method is GET
	if r.Method == http.MethodGet {

		res := AuthResponse{
			Message: "Method GET Not Allowed",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println("Error encoding JSON response:", err)
		}
	}

	// check method is POST
	if r.Method == http.MethodPost {

		var authReq AuthRequestBody

		err := json.NewDecoder(r.Body).Decode(&authReq)
		if err != nil {
			log.Printf("Error decoding auth request: %s\n", err.Error())

			res := AuthResponse{
				Message: "Invalid request payload",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(res)
		}

		if authReq.Username == "admin" && authReq.Password == "admin" {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": authReq.Username,
			})

			tokenString, _ := token.SignedString([]byte("your-256-bit-secret"))

			res := AuthResponse{
				Message: "Authentication successful",
				Token:   tokenString,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(res)

		} else {

			res := AuthResponse{
				Message: "Authentication failed",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(res)

		}

		defer r.Body.Close()

	}

}
