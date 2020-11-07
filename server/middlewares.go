package server

import (
	"github.com/gin-gonic/gin"

	"github.com/OJOMB/git-api-aggregator/log"
)

func (s *Server) middlewares() {
	s.router.Use(s.contentTypeMiddleware())
	s.router.Use(s.logURLServedMiddleWare())

}

func (s *Server) contentTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/" || ctx.Request.URL.Path == "/marco" {
			ctx.Header("Content-Type", "text/html; charset=utf-8")
			log.Info("Content-Type:" + ctx.GetHeader("Content-Type"))
		} else if ctx.Request.URL.Path == "/" {
		} else {
			log.Info("set the content-type to json")
			ctx.Header("Content-Type", "application/json")
		}
	}
}

func (s *Server) logURLServedMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info(ctx.Request.URL.Path + "served")
	}
}
