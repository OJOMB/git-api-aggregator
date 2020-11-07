package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/OJOMB/git-api-aggregator/config"
	healthcontroller "github.com/OJOMB/git-api-aggregator/controllers/health"
	reposcontroller "github.com/OJOMB/git-api-aggregator/controllers/repos"
	userscontroller "github.com/OJOMB/git-api-aggregator/controllers/users"
)

// Server is a server
type Server struct {
	router *gin.Engine
	repos  reposcontroller.ReposControllerInterface
	users  userscontroller.UsersControllerInterface
	health healthcontroller.HealthControllerInterface
	config *config.Config
}

// New returns a new Server instance
func New(router *gin.Engine, config *config.Config) (s *Server) {
	gin.SetMode(gin.ReleaseMode)
	s = &Server{
		router: router,
		repos:  reposcontroller.Repos,
		users:  userscontroller.Users,
		health: healthcontroller.Health,
		config: config,
	}

	// setup middlewares
	s.middlewares()
	// setup all routes
	s.routes()

	return
}

// ListenAndServe For Listening and Serving
func (s *Server) ListenAndServe() {
	addr := fmt.Sprintf("%s:%d", s.config.IP, s.config.Port)
	log.Fatal(s.router.Run(addr))
}
