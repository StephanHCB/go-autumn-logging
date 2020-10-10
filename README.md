# go-autumn-logging

Pluggable logging support for go-autumn.

## About go-autumn

A collection of libraries for [enterprise microservices](https://github.com/StephanHCB/go-mailer-service/blob/master/README.md) in golang that
- is heavily inspired by Spring Boot / Spring Cloud
- is very opinionated
- names modules by what they do
- unlike Spring Boot avoids certain types of auto-magical behaviour
- is not a library monolith, that is every part only depends on the api parts of the other components
  at most, and the api parts do not add any dependencies.  

Fall is my favourite season, so I'm calling it go-autumn.

## About go-autumn-logging

An interface for pluggable logging support with context integration.

This module has zero dependencies, it just provides an interface and a global singleton that is externally
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

If you are writing an application, import one of the modules that actually bring in a logging library, 
such as go-autumn-logging-zerolog. These modules will provide an implementation and place it in the Logger variable.

Of course, you can also provide your own implementation of the `LoggingImplementation` interface, just
set the `Logger` global singleton to an instance of your implementation.

Then just use the Logger, both during application runtime and tests.

### How To Use

Context aware logging is now as simple as:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Logger.Ctx(ctx).Warn("something bad has happened")
}
```

There are convenience functions for those who prefer the function call with arguments style,
though it is a bit less flexible:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Warn(ctx, "something bad has happened")
}
```
