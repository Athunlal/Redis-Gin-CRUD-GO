package jsonOperation

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"
	"github.com/gin-gonic/gin"
)

func ParseUserFromJSON(ctx *gin.Context) (*domain.User, error) {
	user := &domain.User{}
	if err := ctx.BindJSON(&user); err != nil {
		return nil, err
	}
	return user, nil
}
