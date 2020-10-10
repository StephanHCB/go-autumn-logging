package aulogging

import (
	"context"
	"errors"
	"os"
	"testing"
)

// setup / teardown

func TestMain(m *testing.M) {
	beforeTest()
	code := m.Run()
	os.Exit(code)
}

func beforeTest() {
	SetupNoLoggerForTesting()
}

// test cases

func TestTrace(t *testing.T) {
	Trace(context.TODO(), "some message")
}

func TestTracef(t *testing.T) {
	Tracef(context.TODO(), "some format - %v", "some message")
}

func TestTraceErr(t *testing.T) {
	TraceErr(context.TODO(), errors.New("some error"), "some message")
}

func TestTraceErrf(t *testing.T) {
	TraceErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}

func TestDebug(t *testing.T) {
	Debug(context.TODO(), "some message")
}

func TestDebugf(t *testing.T) {
	Debugf(context.TODO(), "some format - %v", "some message")
}

func TestDebugErr(t *testing.T) {
	DebugErr(context.TODO(), errors.New("some error"), "some message")
}

func TestDebugErrf(t *testing.T) {
	DebugErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}

func TestInfo(t *testing.T) {
	Info(context.TODO(), "some message")
}

func TestInfof(t *testing.T) {
	Infof(context.TODO(), "some format - %v", "some message")
}

func TestInfoErr(t *testing.T) {
	InfoErr(context.TODO(), errors.New("some error"), "some message")
}

func TestInfoErrf(t *testing.T) {
	InfoErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}

func TestWarn(t *testing.T) {
	Warn(context.TODO(), "some message")
}

func TestWarnf(t *testing.T) {
	Warnf(context.TODO(), "some format - %v", "some message")
}

func TestWarnErr(t *testing.T) {
	WarnErr(context.TODO(), errors.New("some error"), "some message")
}

func TestWarnErrf(t *testing.T) {
	WarnErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}

func TestError(t *testing.T) {
	Error(context.TODO(), "some message")
}

func TestErrorf(t *testing.T) {
	Errorf(context.TODO(), "some format - %v", "some message")
}

func TestErrorErr(t *testing.T) {
	ErrorErr(context.TODO(), errors.New("some error"), "some message")
}

func TestErrorErrf(t *testing.T) {
	ErrorErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}

func TestFatal(t *testing.T) {
	Fatal(context.TODO(), "some message")
}

func TestFatalf(t *testing.T) {
	Fatalf(context.TODO(), "some format - %v", "some message")
}

func TestFatalErr(t *testing.T) {
	FatalErr(context.TODO(), errors.New("some error"), "some message")
}

func TestFatalErrf(t *testing.T) {
	FatalErrf(context.TODO(), errors.New("some error"), "some format - %v", "some message")
}
