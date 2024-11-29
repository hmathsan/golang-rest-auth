package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "example/rest/test/internal/app/database"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	Message string `json:"message"`
}

func RegisterController(w http.ResponseWriter, r *http.Request) {
	// Ensure Body buffer reader is closed before using
	defer r.Body.Close()

	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if rows, err := db.Db.Query("SELECT * FROM users WHERE username=$1", req.Username); err == nil {
		count := 0
		for rows.Next() {
			count++
		}

		log.Printf("Number of rows found: %d\n", count)
		if count > 0 {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
	} else {
		log.Fatalf("Database error: %v\n", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
	}

	if result, err := db.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", req.Username, req.Password); err == nil {
		id, _ := result.LastInsertId()

		log.Printf("User created with Id: %d\n", id)
		resp := &RegisterUserResponse{Message: fmt.Sprintf("User registered with id %d", id)}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	} else {
		log.Fatalf("Error inserting user to the database: %v\n", err)
		http.Error(w, "Error inserting user to the database", http.StatusInternalServerError)
		return
	}
}
