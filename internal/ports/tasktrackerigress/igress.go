package tasktrackerigress

import (
	"taskTracker/internal/models/tasktracker/dao"
	models "taskTracker/internal/models/tasktracker/tracker"
)

// TaskTrackerAPIPort defines the high-level interface for accessing task-related and user-related API ports.
// It serves as a gateway to obtain specific API ports for tasks and users.
type TaskTrackerAPIPort interface {
	// GetTaskAPIPort returns an instance of TaskAPIPort, which provides methods for managing tasks.
	GetTaskAPIPort() TaskAPIPort
	// GetUserAPIPort returns an instance of UserAPIPort, which provides methods for managing users.
	GetUserAPIPort() UserAPIPort
}

// TaskAPIPort outlines the methods required for managing task-related operations.
// It defines the CRUD (Create, Read, Update, Delete) operations for tasks.
type TaskAPIPort interface {
	// CreateTask inserts a new task into the system.
	// It takes a TaskTracker model as input and returns an error if the operation fails.
	CreateTask(taskRequest models.TaskTracker) error

	// FetchAllTasks retrieves all tasks from the system.
	// It returns a slice of TaskTracker objects and an error if the operation fails.
	FetchAllTasks() ([]dao.TaskTracker, error)

	// UpdateTask modifies an existing task in the system.
	// It takes a map of fields to update and returns an error if the operation fails.
	UpdateTask(updateRequest map[string]interface{}) error

	// DeleteTask removes a task from the system based on its ID.
	// It takes the task ID as a string and returns an error if the operation fails.
	DeleteTask(taskId string) error
}

// UserAPIPort outlines the methods required for managing user-related operations.
// It provides functionality for user creation and validation.
type UserAPIPort interface {
	// CreateUser inserts a new user into the system.
	// It takes a User model containing user details and returns an error if the operation fails.
	CreateUser(user models.User) error

	// ValidateUser checks if the provided user credentials are valid.
	// It takes a User model as input and returns an error if validation fails.
	ValidateUser(user models.User) error
}
