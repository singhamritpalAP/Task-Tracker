package routes

import (
	"github.com/gin-gonic/gin"
	"taskTracker/constants"
	"taskTracker/internal/adapters/ingress/handler/tasktracker"
	"taskTracker/internal/adapters/ingress/handler/user"
	"taskTracker/internal/adapters/ingress/middleware"
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
	taskRouter.setupUserHandler(api)
	return taskRouter.router
}

func (taskRouter *TaskTrackerRouter) setupTaskTrackerHandler(api tasktrackerigress.TaskTrackerAPIPort) {
	handler := tasktracker.NewHandler(api)
	router := taskRouter.router.Group(constants.TaskTrackerGroup)
	{
		router.POST(constants.TaskTrackerEndpoint, middleware.AuthMiddleware(), handler.Create)
		router.GET(constants.TaskTrackerEndpoint, middleware.AuthMiddleware(), handler.FetchAll)
		router.PATCH(constants.TaskTrackerEndpoint, middleware.AuthMiddleware(), handler.Update)
		router.DELETE(constants.TaskTrackerEndpoint, middleware.AuthMiddleware(), handler.Delete)
	}
}

func (taskRouter *TaskTrackerRouter) setupUserHandler(api tasktrackerigress.TaskTrackerAPIPort) {
	handler := user.NewHandler(api)
	router := taskRouter.router.Group(constants.TaskTrackerGroup)
	{
		router.POST(constants.UserEndpoint, handler.Create)
		router.POST(constants.LoginEndpoint, handler.Login)
	}
}
