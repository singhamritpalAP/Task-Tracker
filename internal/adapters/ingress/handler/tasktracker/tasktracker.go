package tasktracker

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taskTracker/constants"
	"taskTracker/internal/adapters/ingress/handler"
	models "taskTracker/internal/models/tasktracker/tracker"
	"taskTracker/internal/models/utils"
	"taskTracker/internal/ports/tasktrackerigress"
)

type Handler struct {
	handler.Interfaces
	api tasktrackerigress.TaskTrackerAPIPort
}

func NewHandler(api tasktrackerigress.TaskTrackerAPIPort) IHandler {
	return &Handler{
		api: api,
	}
}

type IHandler interface {
	handler.Interfaces
}

// Create for handling request to create task
func (handler *Handler) Create(ctx *gin.Context) {
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

	err = handler.api.GetTaskAPIPort().CreateTask(taskRequest)
	if err != nil {
		log.Println("error while creating task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "task created successfully"})
	return
}

// FetchAll to fetch all created tasks for now fetches all tasks in DB.
// can be modified to fetch tasks of logged-in user
func (handler *Handler) FetchAll(ctx *gin.Context) {
	log.Println("Received request for fetch all tasks")
	allTasks, err := handler.api.GetTaskAPIPort().FetchAllTasks()
	if err != nil {
		log.Println("error while fetching all tasks: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// in case there are no tasks
	if len(allTasks) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": allTasks})
	return
}

// Update to update a specific task.
// can be modified to allow a user to update tasks only created by them
func (handler *Handler) Update(ctx *gin.Context) {
	log.Println("Received request for update task")
	var taskUpdateRequest map[string]interface{}
	// fetching update request
	err := ctx.ShouldBindJSON(&taskUpdateRequest)
	if err != nil {
		log.Println(constants.ErrWhileBinding + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrUnableToBindJson.Error()})
		return
	}

	// Validate the status field
	if err := handler.validateStatus(taskUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo validating payload
	//err = taskUpdateRequest.Validate()
	//if err != nil {
	//	log.Println(constants.ErrWhileValidating + err.Error())
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	// passing only updated fields for update
	err = handler.api.GetTaskAPIPort().UpdateTask(taskUpdateRequest)
	if err != nil {
		log.Println("error while updating task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
	return
}

// Delete for deleting a specific task
// can be modified to allow deletion of task created by logged-in user
func (handler *Handler) Delete(ctx *gin.Context) {
	log.Println("Received request for delete task")
	// fetching task id from passed params
	taskId := ctx.Query(constants.TaskIdParam)
	if len(taskId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrTaskIdRequired.Error()})
		return
	}

	err := handler.api.GetTaskAPIPort().DeleteTask(taskId)
	if err != nil {
		log.Println("error while deleting task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
	return
}

// validateStatus checks if the status field is present and valid in the request.
func (handler *Handler) validateStatus(taskUpdateRequest map[string]interface{}) error {
	status, exists := taskUpdateRequest[constants.Status]
	if !exists {
		return nil // case when status is not updated
	}

	// checking if status is string
	statusStr, ok := status.(string)
	if !ok {
		return constants.ErrInvalidStatus
	}

	return utils.IsValidStatus(statusStr)
}
