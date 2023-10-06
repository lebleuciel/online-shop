package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	userRepository "github.com/lebleuciel/online-shop/internal/repository/user"
	userService "github.com/lebleuciel/online-shop/internal/services/user"
)

type Users struct {
	repository *userRepository.UserRepository
	service    *userService.UserService
	// authMiddleware *auth.auth
	authEnabled bool
}

func (u *Users) RegisterRoutes(v1 *gin.RouterGroup) {
	fmt.Println("registering user related endpoints to backend server")
	users := v1.Group("/user")
	users.GET("/", u.getCurrentUser())
}
func (u *Users) getCurrentUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "user data returned")
	}
}
func NewUserModule(userService *userService.UserService, userRepo *userRepository.UserRepository, authEnabled bool) (*Users, error) {
	if userService == nil {
		return nil, ErrNilUserService
	}
	if userRepo == nil {
		return nil, ErrNilUserRepo
	}
	return &Users{
		repository:  userRepo,
		service:     userService,
		authEnabled: authEnabled,
	}, nil
}
