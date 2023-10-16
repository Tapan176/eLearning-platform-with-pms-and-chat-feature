package main

import (
	"log"
	"my-elearning-platform/components/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register API routes from routes/api.go
	routes.RegisterAPIRoutes(r)

	// Define your server settings
	port := ":8080"

	// Start the server
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
