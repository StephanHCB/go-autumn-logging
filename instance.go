package aulogging

import (
	"github.com/StephanHCB/go-autumn-logging/api"
)

// Logger is the singleton instance of ContextAwareLoggingImplementation for use by any
// library and application to access the logger.
//
// See convenience.go for some examples of how to emit log entries.
//
// This global singleton is assigned outside this module. This is so the application author can
// choose which logging library to use.
//
// Library authors:
// only depend on go-autumn-logging and do not touch this variable.
//
// Application authors:
// depend on go-autumn-logging-zerolog (or some other variant) which will set this variable on startup, and
// then you'll be using zerolog.
//
var Logger auloggingapi.LoggingImplementation

// RequestIdRetriever is a singleton instance for use by implementations to obtain
// a tracing request id from a context.
//
// This global singleton is assigned outside this module. Assigning it is optional.
// If an assignment is made, such as by some of the go-autumn tracing implementations,
// then request ids will show up in logs for loggers that support it.
//
// Library authors:
// only depend on go-autumn-logging and do not touch this variable.
//
// Application authors:
// either assign this yourself, or depend on one of the tracing modules.
//
//noinspection GoUnusedGlobalVariable
var RequestIdRetriever auloggingapi.RequestIdRetrieverFunc

// DefaultRequestIdValue lets you override the default value for the request id.
//
// This is useful for console logging if you want your request ids to align.
//
//noinspection GoUnusedGlobalVariable
var DefaultRequestIdValue = ""

// LogEventCallback lets you set up a callback that is called for every log event about to be written
//
//noinspection GoUnusedGlobalVariable
var LogEventCallback auloggingapi.LogEventCallbackFunc
