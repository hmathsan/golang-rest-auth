package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	jwt "example/rest/test/internal/app/auth"
	db "example/rest/test/internal/app/database"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mock user
	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := db.Db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&storedPassword)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Fatalf("Database error: %v\n", err)
		return
	}

	// TODO: hashing
	if password != storedPassword {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJwt(username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
