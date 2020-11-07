package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/OJOMB/git-api-aggregator/config"
)

var (
	// Log it's ya boi
	Log *logrus.Logger
)

// SetupLogging sets up the application logging based on the given config
func SetupLogging(cnfg *config.Config) {
	level, err := logrus.ParseLevel(cnfg.LogLevel)
	if err != nil {
		panic("SETUP LOGGING FAILED: failed to parse log level")
	}
	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}
	if cnfg.Env == "prod" {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}

	Info("Application logger started")
}

// Info for informational logging
func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		// only produce log if the logging level is info or higher
		return
	}
	Log.WithFields(parseFields(tags...))
}

// parseFields takes tags: strings in the format "key:value"
// and returns logrus Field types: alias for map[string]interface{}
func parseFields(tags ...string) (fields logrus.Fields) {
	// when you make a map you can add an optional capacity hint
	// the capacity hint does not limit the size of the map
	// your map will still be able to grow beyond the value of the original capacity hint
	// but it gives the compiler an idea of the the amount of memory to allocate on creation
	fields = make(logrus.Fields, len(tags))

	for _, tag := range tags {
		els := strings.Split(tag, ":")
		fields[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return
}
