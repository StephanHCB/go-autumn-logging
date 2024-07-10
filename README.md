# go-autumn-logging

Pluggable logging support for go-autumn.

## About go-autumn-logging

An interface for pluggable logging support with context integration.

This module has zero dependencies, it just provides an interface and global singletons that are externally
supplied with an implementation of said interface.

Any library or go-autumn module that wants to use context-aware logging should depend on this
module, thereby not pulling in any particular logging implementation.

The application can then pull in the relevant go-autumn-logging submodule, such as go-autumn-logging-zerolog.

## Usage

### Library Authors

If you are writing a library, import **only** this module in your dependencies, none of the other go-autumn-logging-*
modules that actually bring in a logging library.

In your testing code, call `aulogging.SetupNoLoggerForTesting()` to avoid the nil pointer dereference. 

### Application Authors

If you are writing an application, import one of the modules that actually bring in a logging library:

  * [go-autumn-logging-log](https://github.com/StephanHCB/go-autumn-logging-log) using standard log package
  * [go-autumn-slog](https://github.com/roshick/go-autumn-slog) using standard slog package
  * [go-autumn-logging-zerolog](https://github.com/StephanHCB/go-autumn-logging-zerolog) using [zerolog](https://github.com/rs/zerolog)

These modules will provide an implementation and place it in the Logger variable.

Of course, you can also provide your own implementation of the `LoggingImplementation` interface, just
set the `Logger` global singleton to an instance of your implementation.

Then just use the Logger, both during application runtime and tests.

### How To Use

Context aware logging is now as simple as:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Logger.Ctx(ctx).Warn().Print("something bad has happened")
}
```

There are convenience functions for those who prefer the function call with arguments style.
This does the same thing as above:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Warn(ctx, "something bad has happened")
}
```
