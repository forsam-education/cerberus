package errors

import "github.com/forsam-education/cerberus/utils"

// ErrorType is a type of error that can be returned by the proxy.
type ErrorType string

const (
	// ServiceNotFound is the error type when a service cannot be found with the provided path.
	ServiceNotFound ErrorType = "ServiceNotFound"
	// MethodNotAllowed is the error type when a service can't be called with the used HTTP method.
	MethodNotAllowed ErrorType = "MethodNotAllowed"
)

// BuildErrorResponseBody returns a Response correctly formatted with the error status and standard extra data.
func BuildErrorResponseBody(err ErrorType, servicePath string, extraData utils.ResponseExtraData) utils.Response {
	if nil == extraData {
		extraData = utils.ResponseExtraData{}
	}
	extraData["error_type"] = err
	extraData["service_path"] = servicePath
	return utils.Response{
		Status: "error",
		Data:   extraData,
	}
}
