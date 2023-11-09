package repository

import (
	"context"
	"encoding/json"

	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"
	repository_interfaces "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/repository/repository_interface"
	"github.com/go-redis/redis/v8"
)

type User_repo struct {
	Db  *redis.Client
	ctx context.Context
}

func (db *User_repo) Delete(username string) error {
	_, err := db.Db.Del(db.ctx, username).Result()
	if err != nil {
		return err
	}
	return nil
}

func (db *User_repo) Get(username string) (*domain.User, error) {
	ctx := context.Background()

	// Get user data from Redis
	userData, err := db.Db.HGetAll(ctx, username).Result()
	if err != nil {
		return nil, err
	}

	// Unmarshal the user data
	var user domain.User
	if err := json.Unmarshal([]byte(userData["data"]), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update implements repository_interfaces.User_repo_interface.
func (db *User_repo) Update(user *domain.User) error {
	userData, err := user.MarshalBinary()
	if err != nil {
		return err
	}

	err = db.Db.HMSet(db.ctx, user.UserName, "data", userData).Err()
	if err != nil {
		return err
	}
	return nil
}

// Create implements repository_interfaces.User_repo_interface.
func (db *User_repo) Create(user *domain.User) error {
	userData, err := user.MarshalBinary()
	if err != nil {
		return err
	}

	err = db.Db.HMSet(db.ctx, user.UserName, "data", userData).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *redis.Client, Ctx context.Context) repository_interfaces.User_repo_interface {
	return &User_repo{
		Db:  db,
		ctx: Ctx,
	}
}
