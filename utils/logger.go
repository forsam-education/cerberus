package utils

import "github.com/forsam-education/simplelogger"

// Logger is the shared logger for the application.
var Logger simplelogger.Logger

func init() {
	Logger = simplelogger.NewDefaultLogger()
}
