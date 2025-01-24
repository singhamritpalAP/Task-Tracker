package routes

import (
	"github.com/gin-gonic/gin"
	"taskTracker/constants"
	"taskTracker/internal/adapters/ingress/handler/tasktracker"
	"taskTracker/internal/ports/tasktrackerigress"
)

type TaskTrackerRouter struct {
	router *gin.Engine
}

func NewTaskTrackerRouter() *TaskTrackerRouter {
	ginRouter := gin.Default()
	ginRouter.Use(gin.Recovery()) // to recover from any panic
	return &TaskTrackerRouter{
		router: ginRouter,
	}
}

func (taskRouter *TaskTrackerRouter) SetTaskTrackerRoutes(api tasktrackerigress.TaskTrackerAPIPort) *gin.Engine {
	taskRouter.setupTaskTrackerHandler(api)
	return taskRouter.router
}

func (taskRouter *TaskTrackerRouter) setupTaskTrackerHandler(api tasktrackerigress.TaskTrackerAPIPort) {
	handler := tasktracker.NewHandler(api)
	router := taskRouter.router.Group(constants.TaskTrackerGroup)
	{
		router.POST(constants.TaskTrackerEndpoint, handler.Create)
	}
}
