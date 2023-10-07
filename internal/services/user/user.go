package user

import (
	"errors"

	"github.com/gin-gonic/gin"
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
func (u *UserService) GetCurrentUser(c *gin.Context, isAdmin bool) {
	c.JSON(200, "user data returned")
}
