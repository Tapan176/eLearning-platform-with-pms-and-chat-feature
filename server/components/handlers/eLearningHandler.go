package handlers

import (
	"encoding/json"
	"my-elearning-platform/components/models"
	"net/http"
)

// GetCourses retrieves a list of courses
func GetCourses(w http.ResponseWriter, r *http.Request) {
	// Fetch courses from the database (implement this)

	// Simulate courses data for demonstration
	courses := []models.Course{
		{ID: 1, Title: "Course 1", Description: "Description for Course 1"},
		{ID: 2, Title: "Course 2", Description: "Description for Course 2"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}

// Implement other eLearning-related handlers as needed
