package main

import (
	"flag"

	"github.com/OJOMB/git-api-aggregator/app"
)

var env = flag.String("env", "dev", "The environment in which the server is running ['dev', 'test', 'production']")

func main() {
	flag.Parse()
	app.StartApp(*env)
}
