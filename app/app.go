package app

import (
	"github.com/OJOMB/git-api-aggregator/config"
	"github.com/OJOMB/git-api-aggregator/log"
	"github.com/OJOMB/git-api-aggregator/server"
	"github.com/gin-gonic/gin"
)

// StartApp takes some config and handles the initial application startup process
func StartApp(env string) {
	// get the config
	cnfg := config.ConfigMap[env]
	log.SetupLogging(&cnfg)
	// Instantiate application server
	s := server.New(gin.Default(), &cnfg)

	log.Info("Server is starting...", "this:isatag")
	s.ListenAndServe()
}
