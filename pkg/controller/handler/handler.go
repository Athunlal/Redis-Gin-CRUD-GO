package handler

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/controller/handler/handler_interface"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/jsonOperation"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/response"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/usecase/usecase_interface"
	"github.com/gin-gonic/gin"
)

type User_handler struct {
	useCase usecase_interface.User_usecase_interface
}

// Create implements handler_interface.User_handler_interface.
func (us *User_handler) Create(ctx *gin.Context) {

	user, err := jsonOperation.ParseUserFromJSON(ctx)
	if err != nil {
		response.RespondWithError(ctx, 500, err)
		return
	}

	err = us.useCase.Create(user)
	if err != nil {
		response.RespondWithError(ctx, 500, err)
		return
	}

	response.RespondWithSuccess(ctx, "Created successfully")
}

func NewUserhandler(Use usecase_interface.User_usecase_interface) handler_interface.User_handler_interface {
	return &User_handler{
		useCase: Use,
	}
}
