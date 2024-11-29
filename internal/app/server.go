package app

import (
	"log"
	"net/http"

	auth "example/rest/test/internal/app/auth"
	controller "example/rest/test/internal/app/controller"
	handlers "example/rest/test/internal/app/handlers"
	logger "example/rest/test/internal/app/logger"
)

func Server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", controller.LoginController)

	// Protected route
	mux.Handle("/protected", auth.AuthMiddleware(http.HandlerFunc(controller.ProtectedController)))

	// Catch all undefined routes
	mux.Handle("/", http.HandlerFunc(handlers.CatchAllHandler))
	mux.Handle("POST /register", http.HandlerFunc(controller.RegisterController))

	loggerMux := logger.LoggerMiddleware(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggerMux))
}
