package server

import "net/http"

func (s *Server) routes() {
	// serve static files
	s.router.StaticFS("/public/", http.Dir("./public"))
	s.router.StaticFile("/favicon.ico", "./public/favicon-32x32.png")
	s.router.StaticFile("/", "./public/index.html")

	///////////
	// USERS //
	///////////
	s.router.GET("/users/:userName", s.users.GetUser)

	///////////
	// REPOS //
	///////////
	s.router.POST("/users/:userName/repos", s.repos.CreateRepo)

	/////////////
	// HEALTH //
	///////////
	s.router.GET("/polo", s.health.Polo)
}
