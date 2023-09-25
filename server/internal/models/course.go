package models

// Course represents a course model
type Course struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Enrollment    int      `json:"enrollment"`    // Number of students enrolled in the course
	Instructor    string   `json:"instructor"`    // Name of the course instructor
	Duration      int      `json:"duration"`      // Duration of the course in weeks
	Price         int      `json:"price"`         // Course price (in dollars or your currency)
	Level         string   `json:"level"`         // Course level (e.g., "Beginner," "Intermediate," "Advanced")
	Category      string   `json:"category"`      // Course category (e.g., "Programming," "Design," "Business")
	Prerequisites []string `json:"prerequisites"` // List of prerequisite courses or skills
	// Add other course-related fields as needed
}
