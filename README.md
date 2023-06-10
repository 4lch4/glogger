# Glogger

Glogger is my (@4lch4) first attempt at a Go library/module. It's a simple logging library that I use in my projects. It's not meant to be used by anyone else, but if you find it useful, feel free to use it.

## Usage

```go
package main

import (
    "github.com/4lch4/glogger"
)

func main() {
    // Set the application name.
    appName := "AppName"

    // Set the log level to 0 (DEBUG).
    logLevel := 0

    // Create a new logger
    logger := glogger.NewLogger(&glogger.NewLoggerInput{
        AppName:  &appName,
        LogLevel: &logLevel,
    })


    // Log a debug message.
    logger.Debug("This is a debug message.", "main")

    // Log an info message.
    logger.Info("This is an info message.", "main")

    // Log a warning message.
    logger.Warn("This is a warning message.", "main")

    // Log an error message.
    logger.Error("This is an error message.", "main")

    // Log a fatal message.
    logger.Fatal("This is a fatal message.", "main")

    // Log a panic message.
    logger.Panic("This is a panic message.", "main")
}
```
