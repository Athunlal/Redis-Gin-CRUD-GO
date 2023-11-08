package repository

import (
	"fmt"

	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"
	repository_interfaces "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/repository/repository_interface"
	"github.com/go-redis/redis/v8"
)

type User_repo struct {
	Db *redis.Client
}

// Create implements repository_interfaces.User_repo_interface.
func (*User_repo) Create(user *domain.User) error {
	fmt.Println("Created", user)
	return nil
}

func NewRepository(db *redis.Client) repository_interfaces.User_repo_interface {
	return &User_repo{
		Db: db,
	}
}
