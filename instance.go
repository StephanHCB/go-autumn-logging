package aulogging

import "github.com/StephanHCB/go-autumn-logging/api"

// the singleton instance of ContextAwareLoggingImplementation for use by any
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
