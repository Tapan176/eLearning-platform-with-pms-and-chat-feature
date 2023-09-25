package routes

import (
	"github.com/Tapan176/eLearning-platform-with-pms-and-chat-feature/my-elearning-platform/internal/handlers"
	"github.com/gorilla/mux"
)

// RegisterAPIRoutes registers API routes
func RegisterAPIRoutes(r *mux.Router) {
	// Authentication routes
	r.HandleFunc("/auth/login", handlers.Login).Methods("POST")
	r.HandleFunc("/auth/logout", handlers.Logout).Methods("POST")

	// eLearning routes
	r.HandleFunc("/elearning/courses", handlers.GetCourses).Methods("GET")

	// Project management routes
	r.HandleFunc("/projectmanagement/projects", handlers.CreateProject).Methods("POST")

	// Chat routes
	r.HandleFunc("/chat/messages", handlers.SendMessage).Methods("POST")

	// Add more routes as needed
}
