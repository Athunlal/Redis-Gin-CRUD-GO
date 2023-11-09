package controller

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller/handler"
	"github.com/gin-gonic/gin"
)

type User_controller struct {
	handler.User_handler
}

func NewRouter(r *gin.Engine) {
	user := &User_controller{}

	r.POST("/user", user.Create)
}
