package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	auth "example/rest/test/internal/app/auth"
	db "example/rest/test/internal/app/database"
)

type ProtectedResponse struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	UserId   int    `json:"userId"`
}

func ProtectedController(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return auth.SecretKey, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	var userID int
	err := db.Db.QueryRow("SELECT id FROM users WHERE username=$1", username).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	response := &ProtectedResponse{
		Message:  "Protected route authenticated successfully",
		Username: username,
		UserId:   userID,
	}
	json.NewEncoder(w).Encode(response)
}
