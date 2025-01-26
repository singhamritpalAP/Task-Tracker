package tasktraceregress

import "taskTracker/internal/models/tasktracker/dao"

type DbOpsFunc func(ops *DbOps)

type DbOps struct {
	TaskTracker TaskTracker
	User        User
}

// DbPort defines the database operations available for the application.
// It serves as a high-level interface that aggregates different data access layers.
type DbPort interface {
	// GetTaskTracker returns an instance of the TaskTracker interface,
	// which provides methods for managing task tracking operations.
	GetTaskTracker() TaskTracker

	// GetUser returns an instance of the User interface,
	// which provides methods for managing user-related operations.
	GetUser() User
}

// TaskTracker outlines the methods required for managing task tracking operations.
// It defines the CRUD (Create, Read, Update, Delete) operations for task management.
type TaskTracker interface {
	// Create inserts a new task into the database.
	// It takes a TaskTracker object and returns an error if the operation fails.
	Create(task dao.TaskTracker) error

	// Update modifies an existing task in the database.
	// It takes a map of fields to update and optional arguments for conditions,
	// returning an error if the operation fails.
	Update(updateData map[string]interface{}, args ...interface{}) error

	// Read retrieves a single task from the database based on specified criteria.
	// Currently, this method is not implemented and returns an error if called.
	Read() error

	// ReadAll retrieves all tasks from the database.
	// It returns a slice of TaskTracker objects and an error if the operation fails.
	ReadAll() ([]dao.TaskTracker, error)

	// Delete removes a task from the database based on its ID.
	// It takes the task ID as a string and returns an error if the operation fails.
	Delete(taskId string) error
}

// User outlines the methods required for managing user-related operations.
// It provides functionality for user creation and management.
type User interface {
	// Create inserts a new user into the database.
	// It takes a User object containing user details and returns an error if the operation fails.
	Create(userDetails dao.User) error
}
