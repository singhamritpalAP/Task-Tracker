package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taskTracker/constants"
	"taskTracker/core/auth"
	"taskTracker/internal/adapters/ingress/handler"
	models "taskTracker/internal/models/tasktracker/tracker"
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

func (handler *Handler) Create(ctx *gin.Context) {
	log.Println("Received request for creating user")
	var userRequest models.User

	// fetching payload from request
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		log.Println(constants.ErrWhileBinding + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrUnableToBindJson.Error()})
		return
	}

	// todo validating request payload
	//err = userRequest.Validate()
	//if err != nil {
	//	log.Println(constants.ErrWhileValidating + err.Error())
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	err = handler.api.GetUserAPIPort().CreateUser(userRequest)
	if err != nil {
		log.Println("error while creating task: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
	return

}

// Login to handle login request
func (handler *Handler) Login(ctx *gin.Context) {
	var userDetails models.User
	if err := ctx.ShouldBindJSON(&userDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validating user
	err := handler.api.GetUserAPIPort().ValidateUser(userDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For demonstration, for now creating token for userId: 1
	token, err := auth.GenerateJWT(1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
