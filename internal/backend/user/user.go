package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Users struct {
	repository     *repository
	service        *service
	authMiddleware *auth.auth
	authEnabled    bool
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
