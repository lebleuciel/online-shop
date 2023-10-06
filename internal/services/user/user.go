package user

import (
	"errors"

	repository "github.com/lebleuciel/online-shop/internal/repository/user"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) (*UserService, error) {
	if repo == nil {
		return nil, errors.New("user repository can not be nil")
	}
	return &UserService{
		repository: repo,
	}, nil
}
