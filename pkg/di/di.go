package di

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller/handler"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/db"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/repository"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/server"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func InitApi() *gin.Engine {
	client := db.ConnectDB()
	repo := repository.NewRepository(client)
	useCase := usecase.NewUseCase(repo)
	handler.NewUserhandler(useCase)
	route := server.HttpServer()
	controller.NewRouter(route)
	return route
}
