package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	users "github.com/lebleuciel/online-shop/internal/backend/user"
)

type Server struct {
	enviroment string
	engine     *gin.Engine
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.engine.ServeHTTP(w, r)
}
func NewServer(User *users.Users) (*Server, error) {
	if User == nil {
		return nil, errors.New("backend user module can not be nil")
	}
	gin.SetMode("release")
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
	v1 := engine.Group("/api")
	User.RegisterRoutes(v1)
	return &Server{
		enviroment: "release",
		engine:     engine,
	}, nil
}
