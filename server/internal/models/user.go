package models

// User represents a user model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`    // User's email address
	FullName string `json:"fullName"` // User's full name
	Role     string `json:"role"`     // User role (e.g., "student," "instructor," "admin")
	Avatar   string `json:"avatar"`   // URL to the user's profile picture/avatar
	// Add other user-related fields as needed
}
