package tasktrackerservice

import (
	"taskTracker/internal/applications/taskservice"
	"taskTracker/internal/ports/tasktraceregress"
	"taskTracker/internal/ports/tasktrackerigress"
)

type Application struct {
	taskTrackerService tasktrackerigress.TaskAPIPort
}

func NewApplication(database tasktraceregress.DbPort) *Application {
	return &Application{
		taskTrackerService: taskservice.NewApplication(database),
	}
}

func (application *Application) GetTaskAPIPort() tasktrackerigress.TaskAPIPort {
	return application.taskTrackerService
}
