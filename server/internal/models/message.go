package models

// Message represents a chat message model
type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	UserID    int    `json:"userId"`
	Timestamp string `json:"timestamp"` // Timestamp indicating when the message was sent
	IsRead    bool   `json:"isRead"`    // Indicates whether the message has been read by the recipient
	// Add other message-related fields as needed
}
