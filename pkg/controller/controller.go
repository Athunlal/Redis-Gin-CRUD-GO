package controller

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller/handler/handler_interface"
	"github.com/gin-gonic/gin"
)

type User_controller struct {
}

func NewRouter(handler handler_interface.User_handler_interface, r *gin.Engine) {

	r.POST("/user", handler.Create)
}
