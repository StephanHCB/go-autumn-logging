package aulogging

import "context"

// convenience functions for those who prefer this style for readability.
//
// aulogging.Warn(ctx, "some log message")
// aulogging.Warnf(ctx, "some format %v", "what gets put in for the placeholder")
// aulogging.WarnErr(ctx, err, "some log message")

func Trace(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Trace().Print(v...)
	} else {
		Logger.NoCtx().Trace().Print(v...)
	}
}

func Tracef(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Trace().Printf(format, v...)
	} else {
		Logger.NoCtx().Trace().Printf(format, v...)
	}
}

func TraceErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Trace().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Trace().WithErr(err).Print(v...)
	}
}

func TraceErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Trace().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Trace().WithErr(err).Printf(format, v...)
	}
}

func Debug(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Debug().Print(v...)
	} else {
		Logger.NoCtx().Debug().Print(v...)
	}
}

func Debugf(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Debug().Printf(format, v...)
	} else {
		Logger.NoCtx().Debug().Printf(format, v...)
	}
}

func DebugErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Debug().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Debug().WithErr(err).Print(v...)
	}
}

func DebugErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Debug().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Debug().WithErr(err).Printf(format, v...)
	}
}

func Info(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Info().Print(v...)
	} else {
		Logger.NoCtx().Info().Print(v...)
	}
}

func Infof(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Info().Printf(format, v...)
	} else {
		Logger.NoCtx().Info().Printf(format, v...)
	}
}

func InfoErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Info().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Info().WithErr(err).Print(v...)
	}
}

func InfoErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Info().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Info().WithErr(err).Printf(format, v...)
	}
}

func Warn(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Warn().Print(v...)
	} else {
		Logger.NoCtx().Warn().Print(v...)
	}
}

func Warnf(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Warn().Printf(format, v...)
	} else {
		Logger.NoCtx().Warn().Printf(format, v...)
	}
}

func WarnErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Warn().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Warn().WithErr(err).Print(v...)
	}
}

func WarnErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Warn().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Warn().WithErr(err).Printf(format, v...)
	}
}

func Error(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Error().Print(v...)
	} else {
		Logger.NoCtx().Error().Print(v...)
	}
}

func Errorf(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Error().Printf(format, v...)
	} else {
		Logger.NoCtx().Error().Printf(format, v...)
	}
}

func ErrorErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Error().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Error().WithErr(err).Print(v...)
	}
}

func ErrorErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Error().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Error().WithErr(err).Printf(format, v...)
	}
}

func Fatal(ctx context.Context, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Fatal().Print(v...)
	} else {
		Logger.NoCtx().Fatal().Print(v...)
	}
}

func Fatalf(ctx context.Context, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Fatal().Printf(format, v...)
	} else {
		Logger.NoCtx().Fatal().Printf(format, v...)
	}
}

func FatalErr(ctx context.Context, err error, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Fatal().WithErr(err).Print(v...)
	} else {
		Logger.NoCtx().Fatal().WithErr(err).Print(v...)
	}
}

func FatalErrf(ctx context.Context, err error, format string, v ...interface{}) {
	if CtxHasLogger(ctx) {
		Logger.Ctx(ctx).Fatal().WithErr(err).Printf(format, v...)
	} else {
		Logger.NoCtx().Fatal().WithErr(err).Printf(format, v...)
	}
}

// not providing Panic because it is really bad practice to use
