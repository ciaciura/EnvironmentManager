package main

import (
	"EnvManager-api/handlers"
	"EnvManager-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Auth and JWT Middleware
	r.Use(middleware.JWTAuthMiddleware)

	// Server Management Endpoints
	r.HandleFunc("/servers", handlers.GetServers).Methods("GET")
	r.HandleFunc("/servers", handlers.CreateServer).Methods("POST")
	r.HandleFunc("/servers/{id}", handlers.UpdateServer).Methods("PUT")
	r.HandleFunc("/servers/{id}", handlers.DeleteServer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
