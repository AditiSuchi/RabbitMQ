package models

import "time"

// ExportRequest represents a client request to export data
type ExportRequest struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
