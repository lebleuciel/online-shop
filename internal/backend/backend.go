package backend

import (
	"github.com/pkg/errors"

	"github.com/lebleuciel/online-shop/internal/backend/server"
	User "github.com/lebleuciel/online-shop/internal/backend/user"
	UserRepository "github.com/lebleuciel/online-shop/internal/repository/user"
	UserService "github.com/lebleuciel/online-shop/internal/services/user"
)

func NewBackendServer() (*server.Server, error) {

	userRepo, err := UserRepository.NewUserRepository()
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize new user repository")
	}
	userService, err := UserService.NewUserService(userRepo)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize new user service")
	}
	userModule, err := User.NewUserModule(userService, userRepo, true)
	srv, err := server.NewServer(userModule)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize new backend server")
	}
	return srv, nil
}
