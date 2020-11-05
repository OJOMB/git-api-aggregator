package server

import (
	"fmt"
	"log"
	"os"

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

// StartApp takes some config and handles the initial startup process
func StartApp(env string) {
	// get the config
	cnfg := config.ConfigMap[env]

	// get the logger
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")

	// get the router
	router := gin.Default()

	// Instantiate server with shared dependencies
	s := New(router, &cnfg)

	s.ListenAndServe()
}
