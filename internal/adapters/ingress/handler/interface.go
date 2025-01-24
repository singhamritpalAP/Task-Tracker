package handler

import "github.com/gin-gonic/gin"

type Interfaces interface {
	Create(ctx *gin.Context)
}
