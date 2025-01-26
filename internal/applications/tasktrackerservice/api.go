package tasktrackerservice

import (
	"taskTracker/internal/applications/taskservice"
	"taskTracker/internal/applications/userservice"
	"taskTracker/internal/ports/tasktraceregress"
	"taskTracker/internal/ports/tasktrackerigress"
)

type Application struct {
	taskTrackerService tasktrackerigress.TaskAPIPort
	userService        tasktrackerigress.UserAPIPort
}

func NewApplication(database tasktraceregress.DbPort) *Application {
	return &Application{
		taskTrackerService: taskservice.NewApplication(database),
		userService:        userservice.NewApplication(database),
	}
}

func (application *Application) GetTaskAPIPort() tasktrackerigress.TaskAPIPort {
	return application.taskTrackerService
}

func (application *Application) GetUserAPIPort() tasktrackerigress.UserAPIPort {
	return application.userService
}
