package auloggingmock

import (
	"context"
	"github.com/StephanHCB/go-autumn-logging/api"
)

// do-nothing implementation for mocking purposes

// mock for LoggingImplementation

type MockLoggingImplementation struct{}

func (m *MockLoggingImplementation) Ctx(ctx context.Context) auloggingapi.ContextAwareLoggingImplementation {
	return mockContextAware
}

func (m *MockLoggingImplementation) NoCtx() auloggingapi.ContextAwareLoggingImplementation {
	return mockContextAware
}

// mock for ContextAwareLoggingImplementation

var mockContextAware = &MockContextAwareLoggingImplementation{}

type MockContextAwareLoggingImplementation struct{}

func (m *MockContextAwareLoggingImplementation) Trace() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Debug() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Info() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Warn() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Error() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Fatal() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockContextAwareLoggingImplementation) Panic() auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

// mock for ContextAwareLoggingImplementation

var mockLeveled = &MockLeveledLoggingImplementation{}

type MockLeveledLoggingImplementation struct{}

func (m *MockLeveledLoggingImplementation) WithErr(err error) auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockLeveledLoggingImplementation) With(key string, value string) auloggingapi.LeveledLoggingImplementation {
	return mockLeveled
}

func (m *MockLeveledLoggingImplementation) Print(v ...interface{}) {
}

func (m *MockLeveledLoggingImplementation) Printf(format string, v ...interface{}) {
}
