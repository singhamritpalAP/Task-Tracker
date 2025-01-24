package tasktracker

import (
	"github.com/gin-gonic/gin"
	"taskTracker/internal/ports/tasktrackerigress"
)

type TaskTracker struct {
	api tasktrackerigress.TaskTrackerAPIPort
}

func NewHandler(api tasktrackerigress.TaskTrackerAPIPort) *TaskTracker {
	return &TaskTracker{
		api: api,
	}
}
func (taskTracker *TaskTracker) Create(ctx *gin.Context) {
	err := taskTracker.api.GetTaskAPIPort().CreateTask()
	if err != nil {
		// handler error
	}
	return
}
