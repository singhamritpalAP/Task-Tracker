package taskservice

import "taskTracker/internal/ports/tasktraceregress"

type TaskService struct {
	db tasktraceregress.DbPort
}

func NewApplication(database tasktraceregress.DbPort) *TaskService {
	return &TaskService{
		db: database,
	}
}

func (taskService *TaskService) CreateTask() error {
	return nil
}
