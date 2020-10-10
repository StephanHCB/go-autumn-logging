package auloggingmock

import "testing"

func TestNoCtx(t *testing.T) {
	logger := &MockLoggingImplementation{}
	actual := logger.NoCtx()
	if actual != mockContextAware {
		t.FailNow()
	}
}

func TestPanic(t *testing.T) {
	actual := mockContextAware.Panic()
	if actual != mockLeveled {
		t.FailNow()
	}
}

func TestWithKeyValue(t *testing.T) {
	actual := mockLeveled.With("somekey", "some value")
	if actual != mockLeveled {
		t.FailNow()
	}
}
