package glogger

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

// This struct represents an instance of the logger.
type Logger struct {
	// 0 = DEBUG; 1 = INFO; 2 = WARN; 3 = ERROR; 4 = FATAL
	LogLevel int

	// The name of the application using the logger.
	AppName string
}

// The struct used to create a new instance of the Logger struct.
type NewLoggerInput struct {
	AppName  *string
	LogLevel *int
}

// The default value to use for the app name property of the logger if a value is not provided.
var defaultAppName = "Glogger"

// The default value to use for the log level property of the logger if a value is not provided.
var defaultLogLevel = 0

// Get the context for the log message. If the context is empty, the default value of "Glogger" is
// returned.
func (l *Logger) getCtx(ctx string) string {
	if ctx == "" {
		return "Glogger"
	} else {
		return ctx
	}
}

// Gets the log level as a string.
func (l *Logger) getLogLevel() string {
	switch l.LogLevel {
	case 0:
		return "DEBUG"
	case 1:
		return "INFO"
	case 2:
		return "WARN"
	case 3:
		return "ERROR"
	case 4:
		return "FATAL"

	default:
		return "DEBUG"
	}
}

// Gets the log prefix for the log message.
func (l *Logger) getLogPrefix(ctx string) string {
	return "[" + l.AppName + "-" + l.getLogLevel() + "#" + l.getCtx(ctx) + "]: "
}

// Creates a new instance of the Logger struct. Parameters:
//
// - appName: The name of the application using the logger. If nil, the default value is "Glogger".
//
// - logLevel: 0 = Debug / 1 = Info / 2 = Warn / 3 = Error / 4 = Fatal. If nil, the default value is 0.
func NewLogger(loggerInput *NewLoggerInput) *Logger {
	var logLevel *int = loggerInput.LogLevel
	var appName *string = loggerInput.AppName

	if logLevel == nil {
		logLevel = &defaultLogLevel
	}

	if appName == nil {
		appName = &defaultAppName
	}

	return &Logger{
		LogLevel: *logLevel,
		AppName:  *appName,
	}
}

// Outputs a message to the console using the DEBUG log level, fmt.Println, and coloring the
// output bright blue.
func (l *Logger) Debug(msg string, ctx string) {
	if l.LogLevel == 0 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightBlue(logPrefix + msg))
	}
}

// Outputs a message to the console using the DEBUG log level, fmt.Printf, and coloring the
// output bright blue.
func (l *Logger) Debugf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightBlue(msg))
}

// Outputs a message to the console using the INFO log level, fmt.Println, and coloring the
// output cyan.
func (l *Logger) Info(msg string, ctx string) {
	if l.LogLevel <= 1 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.Cyan(logPrefix + "" + msg))
	}
}

// Outputs a message to the console using the INFO log level, fmt.Printf, and coloring the
// output cyan.
func (l *Logger) Infof(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.Cyan(msg))
}

// Outputs a message to the console using the WARN log level, fmt.Println, and coloring the
// output bright yellow.
func (l *Logger) Warn(msg string, ctx string) {
	if l.LogLevel <= 2 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightYellow(logPrefix + msg))
	}
}

// Outputs a message to the console using the WARN log level, fmt.Printf, and coloring the
// output bright yellow.
func (l *Logger) Warnf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightYellow(msg))
}

// Outputs a message to the console using the ERROR log level, fmt.Println, and coloring the
// output bright red.
func (l *Logger) Error(msg string, ctx string) {
	if l.LogLevel <= 3 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

// Outputs a message to the console using the ERROR log level, fmt.Printf, and coloring the
// output bright red.
func (l *Logger) Errorf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightRed(msg))
}

// Outputs a message to the console using the FATAL log level, fmt.Println, and coloring the
// output bright red.
func (l *Logger) Fatal(msg string, ctx string) {
	if l.LogLevel <= 4 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

// Outputs a message to the console using the FATAL log level, fmt.Printf, and coloring the
// output bright red.
func (l *Logger) Fatalf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightRed(msg))
}

// Outputs a message to the console using the SUCCESS log level, fmt.Println, and coloring the
// output bright green.
func (l *Logger) Success(msg string, ctx string) {
	logPrefix := l.getLogPrefix(ctx)
	fmt.Println(aurora.BrightGreen(logPrefix + msg))
}

// Outputs a message to the console using the SUCCESS log level, fmt.Printf, and coloring the
// output bright green.
func (l *Logger) Successf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightGreen(msg))
}
