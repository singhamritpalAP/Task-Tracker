package taskservice

import (
	"log"
	"taskTracker/constants"
	"taskTracker/internal/models/tasktracker/dao"
	models "taskTracker/internal/models/tasktracker/tracker"
	"taskTracker/internal/models/utils"
	"taskTracker/internal/ports/tasktraceregress"
)

// TaskService provides methods for managing task-related operations.
// It interacts with the database through the DbPort interface.
type TaskService struct {
	db tasktraceregress.DbPort // Interface for database operations
}

// NewApplication initializes a new TaskService instance with the provided database interface.
// It returns a pointer to the TaskService.
func NewApplication(database tasktraceregress.DbPort) *TaskService {
	return &TaskService{
		db: database,
	}
}

// CreateTask creates a new task in the database.
// It takes a TaskTracker model as input, converts it to a DAO, and stores it in the database.
func (taskService *TaskService) CreateTask(taskRequest models.TaskTracker) error {
	log.Println("Creating task: ", taskRequest)
	// Convert the TaskTracker model to its DAO representation for storage
	taskDao := utils.ConvertTaskToDao(taskRequest)

	// Call the database function to create the task
	err := taskService.db.GetTaskTracker().Create(taskDao)
	if err != nil {
		log.Println("Error creating task: ", err)
		return err
	}
	log.Println("Task created")
	return nil // Return nil indicating success
}

// FetchAllTasks retrieves all tasks from the database.
// It returns a slice of TaskTracker objects and an error if the operation fails.
func (taskService *TaskService) FetchAllTasks() ([]dao.TaskTracker, error) {
	log.Println("Fetching all tasks")
	// Call the database function to fetch all tasks
	tasks, err := taskService.db.GetTaskTracker().ReadAll()
	if err != nil {
		log.Println("Error fetching tasks: ", err)
		return nil, err
	}
	return tasks, nil
}

// UpdateTask modifies an existing task in the database.
// It takes a map of fields to update and returns an error if the operation fails.
func (taskService *TaskService) UpdateTask(updateRequest map[string]interface{}) error {
	taskId := updateRequest[constants.TaskId] // Retrieve task ID from the update request
	if taskId == nil {
		return constants.ErrTaskIdRequired // Return an error if task ID is not provided
	}
	log.Println("Updating task with id: ", taskId)

	// Call the database function to update the task
	return taskService.db.GetTaskTracker().Update(updateRequest, taskId)
}

// DeleteTask removes a task from the database based on its ID.
// It takes the task ID as a string and returns an error if the operation fails.
func (taskService *TaskService) DeleteTask(taskId string) error {
	log.Println("Deleting task with id: ", taskId)
	// Call the database function to delete the task
	return taskService.db.GetTaskTracker().Delete(taskId)
}
