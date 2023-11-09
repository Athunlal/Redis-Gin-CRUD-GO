package db

import "github.com/go-redis/redis/v8"

func ConnectDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})
	return client
}
