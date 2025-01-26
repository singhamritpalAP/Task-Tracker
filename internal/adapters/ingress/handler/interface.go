package handler

import "github.com/gin-gonic/gin"

type Interfaces interface {
	Create(ctx *gin.Context)
	FetchAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Login(ctx *gin.Context)
}
