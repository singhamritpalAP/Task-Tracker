package taskservice

import (
	"log"
	models "taskTracker/internal/models/tasktracker/tracker"
	"taskTracker/internal/ports/tasktraceregress"
)

type TaskService struct {
	db tasktraceregress.DbPort
}

func NewApplication(database tasktraceregress.DbPort) *TaskService {
	return &TaskService{
		db: database,
	}
}

func (taskService *TaskService) CreateTask(taskRequest models.TaskTracker) error {
	log.Println("Creating task")
	// call db function to save task
	return nil
}

func (taskService *TaskService) FetchAllTasks() ([]models.TaskTracker, error) {
	// call db function to fetch all tasks
	return []models.TaskTracker{}, nil
}

func (taskService *TaskService) UpdateTask(updateRequest models.TaskUpdate) error {
	log.Println("Updating task with id: ", updateRequest.TaskId)
	//  call db function to update task
	return nil
}

func (taskService *TaskService) DeleteTask(taskId string) error {
	log.Println("Deleting task with id: ", taskId)
	// call db function to delete task
	return nil
}
