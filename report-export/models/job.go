package models

import "time"

// JobStatus represents the state of the export job
type JobStatus string

const (
	StatusPending   JobStatus = "PENDING"
	StatusRunning   JobStatus = "RUNNING"
	StatusCompleted JobStatus = "COMPLETED"
	StatusFailed    JobStatus = "FAILED"
)

// ExportJob represents a background job processing the export request
type ExportJob struct {
	ID        string    `json:"id"`
	RequestID string    `json:"request_id"`
	Status    JobStatus `json:"status"`
	ResultURL string    `json:"result_url,omitempty"`
	Error     string    `json:"error,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
