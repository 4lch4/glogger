# Glogger

Glogger is my (@4lch4) first attempt at a Go library/module. It's a simple logging library that I use in my projects. It's not meant to be used by anyone else, but if you find it useful, feel free to use it.

## Usage

```go
package main

import (
    "github.com/4lch4/Glogger"
)

func main() {
    // Create a new logger
    logger := glogger.NewLogger()

    // Set the log level
    logger.SetLevel(glogger.DEBUG)

    // Log a message
    logger.Log(glogger.DEBUG, "This is a debug message")

    // Log a message with a format
    logger.Logf(glogger.INFO, "This is a %s message", "info")

    // Log a message with a format and a variable number of arguments
    logger.Logvf(glogger.WARN, "This is a %s message with %d arguments", "warning", 2, "foo", "bar")

    // Log a message with a format and a slice of arguments
    logger.Logvsf(glogger.ERROR, "This is a %s message with %d arguments", "error", 2, []interface{}{"foo", "bar"})

    // Log a message with a format and a map of arguments
    logger.Logvmf(glogger.FATAL, "This is a %s message with %d arguments", "fatal", 2, map[string]interface{}{"foo": "bar", "baz": "qux"})
}
```
