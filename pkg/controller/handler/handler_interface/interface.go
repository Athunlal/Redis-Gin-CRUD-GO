package handler_interface

import (
	"github.com/gin-gonic/gin"
)

type User_handler_interface interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
