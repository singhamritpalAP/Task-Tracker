package dao

import (
	"log"
	"taskTracker/constants"
	"taskTracker/core/relationaldatabase"
	"taskTracker/internal/models/tasktracker/dao"
)

// TrackerAdapter serves as an adapter for task tracking operations,
// providing a layer of abstraction over the database interactions.
type TrackerAdapter struct {
	wrapper relationaldatabase.RelationalDatabase // Database wrapper for executing queries
}

// NewTrackerAdapter initializes a new TrackerAdapter with the provided database wrapper.
func NewTrackerAdapter(wrapper *relationaldatabase.DbWrapper) *TrackerAdapter {
	return &TrackerAdapter{wrapper: wrapper}
}

// Create inserts a new task into the database.
// It takes a TaskTracker object and returns an error if the operation fails.
func (tracker *TrackerAdapter) Create(task dao.TaskTracker) error {
	return tracker.wrapper.Create(&dao.TaskTracker{}, &task)
}

// Update modifies an existing task in the database.
// It takes a map of fields to update and optional arguments for conditions.
func (tracker *TrackerAdapter) Update(updateData map[string]interface{}, args ...interface{}) error {
	return tracker.wrapper.Update(&dao.TaskTracker{}, updateData, constants.WhereId, args)
}

// Read retrieves a single task from the database based on specified criteria.
// Currently, this method is not implemented and returns nil.
func (tracker *TrackerAdapter) Read() error {
	return nil
}

// ReadAll retrieves all tasks from the database.
// It returns a slice of TaskTracker objects and an error if the operation fails.
func (tracker *TrackerAdapter) ReadAll() ([]dao.TaskTracker, error) {
	var tasks []dao.TaskTracker            // Slice to hold retrieved tasks
	err := tracker.wrapper.ReadAll(&tasks) // Call to the wrapper to read all tasks
	if err != nil {
		log.Println("Error retrieving tasks:", err)
		return nil, err
	}
	return tasks, nil // Return the slice of tasks
}

// Delete removes a task from the database based on its ID.
// It takes the task ID as a string and returns an error if the operation fails.
func (tracker *TrackerAdapter) Delete(taskId string) error {
	return tracker.wrapper.Delete(&dao.TaskTracker{}, taskId, constants.WhereId, taskId)
}
