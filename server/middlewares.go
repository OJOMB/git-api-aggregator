package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (s *Server) middlewares() {
	s.router.Use(s.contentTypeMiddleware())
	s.router.Use(s.logURLServedMiddleWare())

}

func (s *Server) contentTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/" {
			ctx.Header("Content-Type", "text/html; charset=utf-8")
			log.Printf("Content-Type: %s", ctx.GetHeader("Content-Type"))
		} else if ctx.Request.URL.Path == "/" {
		} else {
			log.Print("set the content-type to json")
			ctx.Header("Content-Type", "application/json")
		}
	}
}

func (s *Server) logURLServedMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("%s served", ctx.Request.URL.Path)
	}
}
