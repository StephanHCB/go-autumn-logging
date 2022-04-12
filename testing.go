package aulogging

import "github.com/StephanHCB/go-autumn-logging/mock"

// setup function to use in tests that injects a do-nothing logger. You will need this if you are writing a library.
//
// see convenience_test.go for an example how to use this.
func SetupNoLoggerForTesting() {
	Logger = &auloggingmock.MockLoggingImplementation{}
}
