package repository_interfaces

import "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"

type User_repo_interface interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	Get(username string) (*domain.User, error)
	Delete(username string) error
}
