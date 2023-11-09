package controller

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller/handler/handler_interface"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler handler_interface.User_handler_interface, r *gin.Engine) {

	r.POST("/user", handler.Create)
	r.PUT("/user", handler.Update)
	r.GET("/user", handler.Get)
	r.DELETE("/user", handler.Delete)
}
