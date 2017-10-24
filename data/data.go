package data

import (
	"encoding/json"
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

// ToJSON takes a slice of tasks and returns their JSON representation.
func ToJSON(tasks []Task) (tasksJSON []byte, err error) {
	return json.Marshal(tasks)
}

// ToJSON returns the JSON representation of the given task.
func (task Task) ToJSON() (taskJSON []byte, err error) {
	return json.Marshal(task)
}

// GetTasksDummy returns a slice of tasks filled with static dummy data.
func GetTasksDummy() []Task {

	task1 := Task{
		ID:    1,
		Title: "The first task ever created!",
		Content: `
		# Headline
		
		Content of a task:
		* Follows the [CommonMark](http://commonmark.org/) spec
		* Render a string as markdown
		`,
		TimeCreated: time.Now(),
		SubTasks: []Task{
			{
				ID:    2,
				Title: "This is the first subtask of task 1!",
				Content: `
			  # Headline
			
			  Content of a task:
			  * Follows the [CommonMark](http://commonmark.org/) spec
			  * Render a string as markdown
			`,
				TimeCreated: time.Now(),
				SubTasks: []Task{
					{
						ID:    3,
						Title: "Hit is the first subtask of subtask 2!",
						Content: `
						# Headline
						
						Content of a task:
						* Follows the [CommonMark](http://commonmark.org/) spec
						* Render a string as markdown
						`,
						TimeCreated: time.Now(),
					},
				},
			},
		},
		Updates: []TaskUpdate{
			{
				ID:          1,
				TimeCreated: time.Now(),
				Type:        PhoneCall,
				Title:       "Update 1: Call from Mr. X",
				Description: `
				# Headline
				
				Content of a task:
				* Follows the [CommonMark](http://commonmark.org/) spec
				* Render a string as markdown
				`,
			},
			{
				ID:          2,
				TimeCreated: time.Now(),
				Type:        File,
				Title:       "Update 2: Last years sales",
				Description: `
				# Headline
				
				Content of a task:
				* Follows the [CommonMark](http://commonmark.org/) spec
				* Render a string as markdown
				`,
			},
			{
				ID:          3,
				TimeCreated: time.Now(),
				Type:        Mail,
				Title:       "Update 3: Mail from Mr. X",
				Description: `
				# Headline
				
				Content of a task:
				* Follows the [CommonMark](http://commonmark.org/) spec
				* Render a string as markdown
				`,
			},
		},
	}

	return []Task{task1}
}
