package utils

import (
	"github.com/forsam-education/simplelogger"
)

// Logger is the shared logger for the application.
var Logger simplelogger.Logger

// VerboseFlag describes if Cerberus is in verbose mode.
var VerboseFlag bool

func init() {
	Logger = simplelogger.NewDefaultLogger(simplelogger.DEBUG)
}

// LogVerbose is used to log only if verbose mode is enabled.
func LogVerbose(message string, data map[string]interface{}) {
	if VerboseFlag {
		Logger.Debug(message, data)
	}
}
