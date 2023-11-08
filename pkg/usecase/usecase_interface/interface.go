package usecase_interface

import "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"

type User_usecase_interface interface {
	Create(user *domain.User) error
}
