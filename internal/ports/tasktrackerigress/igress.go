package tasktrackerigress

import models "taskTracker/internal/models/tasktracker/tracker"

type TaskTrackerAPIPort interface {
	GetTaskAPIPort() TaskAPIPort
}

type TaskAPIPort interface {
	CreateTask(taskRequest models.TaskTracker) error
	FetchAllTasks() ([]models.TaskTracker, error)
	UpdateTask(updateRequest models.TaskUpdate) error
	DeleteTask(taskId string) error
}
