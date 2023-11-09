package server

import "github.com/gin-gonic/gin"

func HttpServer() *gin.Engine {
	r := gin.Default()
	return r
}
