package models

// Project represents a project model
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"` // Start date of the project
	EndDate     string `json:"endDate"`   // End date of the project
	Status      string `json:"status"`    // Status of the project (e.g., "In Progress", "Completed", etc.)
	OwnerID     int    `json:"ownerId"`   // ID of the user who owns the project
	// Add other project-related fields as needed
}
