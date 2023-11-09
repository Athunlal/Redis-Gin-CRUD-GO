// user_repo_test.go

package repository

import (
	"context"
	"testing"

	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

// Define an interface for redis.Client methods you use
type RedisClient interface {
	HMSet(ctx context.Context, key string, fields ...interface{}) *redis.StatusCmd
	// Add other methods as needed
}

// MockDB is a mock implementation of the RedisClient interface
type MockDB struct {
	mock.Mock
}

func (m *MockDB) HMSet(ctx context.Context, key string, fields ...interface{}) *redis.StatusCmd {
	args := m.Called(ctx, key, fields...)
	return args.Get(0).(*redis.StatusCmd)
	// Implement other mocked methods similarly
}

func TestCreateUser(t *testing.T) {
	// Set up the mock database
	mockDB := new(MockDB)

	// Create an instance of User_repo with the mock database
	userRepo := User_repo{
		Db:  mockDB,
		ctx: context.TODO(), // Provide the context as needed
	}

	// Create a test user
	testUser := &domain.User{
		UserName: "testuser",
		// Set other fields as needed
	}

	// Mock the expected calls
	mockDB.On("HMSet", mock.Anything, "testuser", mock.Anything).Return(redis.NewStatusCmd())

	// Call the Create method
	err := userRepo.Create(testUser)

	// Assert that the mock was called as expected
	mockDB.AssertExpectations(t)

	// Check if the result is as expected
	if err != nil {
		t.Errorf("Unexpected error while creating user: %v", err)
	}
}
