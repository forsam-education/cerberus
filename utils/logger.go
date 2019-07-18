package utils

import (
	"github.com/forsam-education/simplelogger"
	"os"
)

// Logger is the shared logger for the application.
var Logger simplelogger.Logger

func init() {
	Logger = simplelogger.NewDefaultLogger()
}

// LogAndForceExit is a quick helper to force exit on unrecoverable errors.
func LogAndForceExit(err error) {
	Logger.StdError(err, nil)
	os.Exit(1)
}
