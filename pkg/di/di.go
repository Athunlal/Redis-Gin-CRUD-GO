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
	client, ctx := db.ConnectDB()
	repo := repository.NewRepository(client, ctx)
	useCase := usecase.NewUseCase(repo)
	handler := handler.NewUserhandler(useCase)
	route := server.HttpServer()
	controller.NewRouter(handler, route)
	return route
}
