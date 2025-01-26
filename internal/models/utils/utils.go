package utils

import (
	"taskTracker/internal/models/tasktracker/dao"
	models "taskTracker/internal/models/tasktracker/tracker"
)

func ConvertTaskToDao(task models.TaskTracker) dao.TaskTracker {
	return dao.TaskTracker{
		TaskId:      task.TaskId,
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
