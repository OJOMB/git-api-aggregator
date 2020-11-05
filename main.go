package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// Config holds the configuration settings for an instance of the app
type Config struct {
	Env  string
	IP   string
	Port int
}

func getGithubAccessToken() string {
	return githubAccessToken
}
