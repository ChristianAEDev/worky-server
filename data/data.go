package data

import (
	"time"
)

// Task represents a piece of work. A task can have multiple sub tasks
type Task struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	TimeCreated time.Time    `json:"timeCreated"`
	SubTasks    []Task       `json:"subTasks"`
	Updates     []TaskUpdate `json:"updated"`
}

// TaskUpdate represents a single piece of information that adds data to a task.
type TaskUpdate struct {
	ID          int            `json:"id"`
	TimeCreated time.Time      `json:"timeCreated"`
	Type        TaskUpdateType `json:"type"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
}

// TaskUpdateType represents the kind of data a TaskUpdate represents.
type TaskUpdateType string

const (
	// PhoneCall represents an interaction on the telephone
	PhoneCall TaskUpdateType = "PHONE_CALL"
	// File represents a file
	File = "FILE"
	// Mail represents an email
	Mail = "MAIL"
)
