package auloggingapi

import "context"

// LoggingImplementation is the top level entry point to our logging API.
//
// Note: access to the singleton instance is provided in the Logger global variable.
//
// See convenience.go for some examples of how to use these interfaces.
type LoggingImplementation interface {
	// Ctx returns a context aware logger.
	//
	// You should use this if you have a context.
	//
	// Behaviour is implementation dependent, usually the logger will be taken from the context.
	Ctx(ctx context.Context) ContextAwareLoggingImplementation

	// NoCtx returns a context aware logger using a blank context.
	//
	// Use this if you do not have a context (but if it's your own code, you really should).
	NoCtx() ContextAwareLoggingImplementation
}

type ContextAwareLoggingImplementation interface {
	// Trace logs using TRACE level (-1).
	Trace() LeveledLoggingImplementation

	// Debug logs using DEBUG level (0).
	Debug() LeveledLoggingImplementation

	// Info logs using INFO level (1).
	Info() LeveledLoggingImplementation

	// Warn logs using WARN level (2).
	Warn() LeveledLoggingImplementation

	// Error logs using ERROR level (3).
	Error() LeveledLoggingImplementation

	// Fatal logs using FATAL level (4) and exit the application.
	//
	// Note that once application startup is over, this is usually very bad practice.
	Fatal() LeveledLoggingImplementation

	// Panic logs using PANIC level (5), print a stack trace, and exit the application.
	//
	// Note that this is usually very bad practice.
	Panic() LeveledLoggingImplementation
}

type LeveledLoggingImplementation interface {
	// WithErr adds an error to the log message and keeps going.
	//
	// This does not actually log anything, you need to follow it up by Print or Printf.
	WithErr(err error) LeveledLoggingImplementation

	// With adds an additional key-value-pair to the log message and keeps going.
	//
	// This does not actually log anything, you need to follow it up by Print or Printf.
	With(key string, value string) LeveledLoggingImplementation

	// Print finalizes the log entry and emits it, including all arguments in the message.
	Print(v ...interface{})

	// Printf finalizes the log entry and emits it, using a format string for the message.
	Printf(format string, v ...interface{})
}

// RequestIdRetrieverFunc is the type of a function used to obtain a request id from a context
//
// Note: access to the singleton instance is provided in the Logger global variable.
//
// used for tracing integration
type RequestIdRetrieverFunc func(ctx context.Context) string

// LogEventCallbackFunc is the type of a callback function you can provide to receive
// a callback for every log event.
//
// level is set to "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"
type LogEventCallbackFunc func(ctx context.Context, level string, message string, err error, additionalFields map[string]string)

// CtxHasLoggerFunc is the type of a function that checks whether a context.Context
// contains a logger.
type CtxHasLoggerFunc func(ctx context.Context) bool
