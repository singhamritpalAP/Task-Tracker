package tasktracker

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taskTracker/constants"
	models "taskTracker/internal/models/tasktracker/tracker"
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
	log.Println("Received request for create task")
	var taskRequest models.TaskTracker

	// fetching payload from request
	err := ctx.ShouldBindJSON(&taskRequest)
	if err != nil {
		log.Println(constants.ErrWhileBinding + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrUnableToBindJson.Error()})
		return
	}

	// validating request payload
	err = taskRequest.Validate()
	if err != nil {
		log.Println(constants.ErrWhileValidating + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = taskTracker.api.GetTaskAPIPort().CreateTask(taskRequest)
	if err != nil {
		log.Println("error while creating task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "task created successfully"})
	return
}

func (taskTracker *TaskTracker) FetchAll(ctx *gin.Context) {
	log.Println("Received request for fetch all tasks")
	allTasks, err := taskTracker.api.GetTaskAPIPort().FetchAllTasks()
	if err != nil {
		log.Println("error while fetching all tasks: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(allTasks) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "no tasks found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": allTasks})
	return
}

func (taskTracker *TaskTracker) Update(ctx *gin.Context) {
	log.Println("Received request for update task")
	var taskUpdateRequest models.TaskUpdate
	// fetching update request
	err := ctx.ShouldBindJSON(&taskUpdateRequest)
	if err != nil {
		log.Println(constants.ErrWhileBinding + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrUnableToBindJson.Error()})
		return
	}

	// validating payload
	err = taskUpdateRequest.Validate()
	if err != nil {
		log.Println(constants.ErrWhileValidating + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = taskTracker.api.GetTaskAPIPort().UpdateTask(taskUpdateRequest)
	if err != nil {
		log.Println("error while updating task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
	return
}

func (taskTracker *TaskTracker) Delete(ctx *gin.Context) {
	log.Println("Received request for delete task")
	taskId := ctx.Query("taskId")
	if len(taskId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrTaskIdRequired.Error()})
		return
	}

	err := taskTracker.api.GetTaskAPIPort().DeleteTask(taskId)
	if err != nil {
		log.Println("error while deleting task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
	return
}
