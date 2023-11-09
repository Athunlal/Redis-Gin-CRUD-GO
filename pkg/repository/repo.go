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

// Update implements repository_interfaces.User_repo_interface.
func (db *User_repo) Update(user *domain.User) error {

	// Marshal user data
	userData, err := user.MarshalBinary()
	if err != nil {
		return err
	}

	// Update the user data in Redis with a key based on the username
	err = db.Db.HMSet(db.Dbcontext(), user.UserName, "data", userData).Err()
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}
	fmt.Println("User updated:", user.UserName)
	return nil
}

// Create implements repository_interfaces.User_repo_interface.
func (db *User_repo) Create(user *domain.User) error {

	userData, err := user.MarshalBinary()
	if err != nil {
		return err
	}

	err = db.Db.HMSet(db.Dbcontext(), user.UserName, "data", userData).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *redis.Client) repository_interfaces.User_repo_interface {
	return &User_repo{
		Db: db,
	}
}
