package user

import "errors"

type UserService struct {
	repository *repository
}

func NewUserService(repo *repository) (*UserService, error) {
	if repo == nil {
		return nil, errors.New("user repository can not be nil")
	}
	return &UserService{
		repository: repo,
	}, nil
}
