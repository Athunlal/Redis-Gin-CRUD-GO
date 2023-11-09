package usecase

import (
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/domain"
	repository_interfaces "github.com/athunlal/Redis-Gin-CRUD-Go/pkg/repository/repository_interface"
	"github.com/athunlal/Redis-Gin-CRUD-Go/pkg/usecase/usecase_interface"
)

type User_usecase struct {
	Repo repository_interfaces.User_repo_interface
}

// Create implements usecase_interface.User_usecase_interface.
func (r *User_usecase) Create(user *domain.User) error {
	if err := r.Repo.Create(user); err != nil {
		return err
	}

	return nil
}

func NewUseCase(repo repository_interfaces.User_repo_interface) usecase_interface.User_usecase_interface {
	return &User_usecase{
		Repo: repo,
	}
}
