package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken returns the secret Github access token
func GetGithubAccessToken() string {
	return githubAccessToken
}

// Config holds all application configuration values
type Config struct {
	Env      string
	IP       string
	Port     uint
	App      string
	Static   string
	LogLevel string
}

// ConfigMap holds the configuration data for each given environment
var ConfigMap = map[string]Config{
	"dev": {
		Env:      "dev",
		IP:       "0.0.0.0",
		Port:     8080,
		App:      "mvsc-microservice",
		Static:   "./public",
		LogLevel: "info",
	},
}
