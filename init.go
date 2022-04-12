package aulogging

import auloggingmock "github.com/StephanHCB/go-autumn-logging/mock"

// init ensures no nil pointer exception will be thrown (though no logging output will be produced and no callbacks will be made)
//
// Application authors should import one of the provided implementation packages to actually get logging.
func init() {
	if Logger == nil {
		Logger = &auloggingmock.MockLoggingImplementation{}
	}
}
