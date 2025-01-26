package utils

import (
	"taskTracker/constants"
	"taskTracker/internal/models/tasktracker/dao"
	models "taskTracker/internal/models/tasktracker/tracker"
)

func ConvertTaskToDao(task models.TaskTracker) dao.TaskTracker {
	return dao.TaskTracker{
		// TaskId:      task.TaskId,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
	}
}

func ConvertUserToDao(user models.User) dao.User {
	return dao.User{
		Username: user.Username,
		Password: user.Password,
	}
}

// IsValidStatus checks if the provided status is a valid TaskStatus.
// It returns an error if the status is invalid.
func IsValidStatus(status string) error {
	switch models.TaskStatus(status) {
	case models.StatusStarted, models.StatusPending, models.StatusCompleted:
		return nil // Valid status
	default:
		return constants.ErrInvalidStatus // Return an error for invalid status
	}
}
