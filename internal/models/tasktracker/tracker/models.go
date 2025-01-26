package models

import (
	"taskTracker/constants"
)

// TaskStatus a custom type for Task Status
type TaskStatus string

// valid statuses as constants
const (
	StatusStarted   TaskStatus = constants.Started
	StatusPending   TaskStatus = constants.Pending
	StatusCompleted TaskStatus = constants.Completed
)

// TaskTracker represents a task with its details.
type TaskTracker struct {
	TaskId      string     `json:"task_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	//StartDate   time.Time  `json:"start_date"`
	//EndDate     time.Time  `json:"end_date"`
}

// TaskUpdate represents the fields that can be updated for a task.
type TaskUpdate struct {
	TaskId      string     `json:"task_id"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty"`
}

// User represents the fields for user todo add validation
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Store hashed passwords
}

// Validate to validate received payload
func (task *TaskTracker) Validate() error {
	if task.TaskId == "" {
		return constants.ErrTaskIdRequired
	}
	if len(task.Title) < 2 {
		return constants.ErrTitleRequired
	}
	if task.Description == "" {
		return constants.ErrDescriptionRequired
	}
	if !isValidStatus(task.Status) {
		return constants.ErrInvalidStatus
	}
	//if task.StartDate.IsZero() {
	//	return constants.ErrStartDateRequired
	//}
	//if task.EndDate.Before(task.StartDate) {
	//	return constants.ErrEndDateRequired
	//}

	return nil
}

func (taskUpdate *TaskUpdate) Validate() error {
	if taskUpdate.TaskId == "" {
		return constants.ErrTaskIdRequired
	}

	// if task title being updated it must have at least 2 chars
	if taskUpdate.Title != "" && len(taskUpdate.Title) < 2 {
		return constants.ErrTitleRequired
	}
	if taskUpdate.Description != "" && len(taskUpdate.Description) < 2 {
		return constants.ErrDescriptionRequired
	}
	if taskUpdate.Status != "" && !isValidStatus(taskUpdate.Status) {
		return constants.ErrInvalidStatus
	}
	return nil
}

// isValidStatus checks if the provided status is valid.
func isValidStatus(status TaskStatus) bool {
	switch status {
	case StatusStarted, StatusPending, StatusCompleted:
		return true
	default:
		return false
	}
}
