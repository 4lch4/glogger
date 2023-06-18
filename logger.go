package glogger

import (
	"fmt"
	"path"
	"runtime"

	"github.com/logrusorgru/aurora/v4"
)

type (
	NewLoggerInput struct {
		// The name of the application using the logger.
		AppName *string

		// 0 = DEBUG; 1 = INFO; 2 = WARN; 3 = ERROR; 4 = FATAL
		LogLevel *int
	}

	Logger struct {
		// The name of the application using the logger.
		AppName string

		// 0 = DEBUG; 1 = INFO; 2 = WARN; 3 = ERROR; 4 = FATAL
		LogLevel int
	}
)

// #region Internals

// A map of log levels with the key as a string of the log level name or the log level number.
// This means that the following are valid keys:
//
// - "DEBUG" | "0"
//
// - "INFO" | "1"
//
// - "WARNING" | "2"
//
// - "ERROR" | "3"
//
// - "FATAL" | "4"
var logLevels = map[string]string{
	"DEBUG":   "0",
	"INFO":    "1",
	"WARNING": "2",
	"ERROR":   "3",
	"FATAL":   "4",
	"0":       "DEBUG",
	"1":       "INFO",
	"2":       "WARNING",
	"3":       "ERROR",
	"4":       "FATAL",
}

// The default value to use for the app name property of the logger if a value is not provided.
var defaultAppName = "Glogger"

// The default value to use for the log level property of the logger if a value is not provided.
var defaultLogLevel = 0

// Get the context for the log message. If the context is empty, the default value of "Glogger" is
// returned.
func (l *Logger) getCtx(ctx *string) string {
	if ctx == nil {
		return defaultAppName
	}

	return *ctx
}

// Gets the log prefix for the log message.
func (l *Logger) getLogPrefix(logLevel *int, ctx *string) string {
	if logLevel == nil {
		logLevel = &l.LogLevel
	}

	return "[" + l.AppName + "-" + l.ConvertLogLevel(logLevel) + "#" + l.getCtx(ctx) + "]: "
}

// #endregion Internals

// #region Exported

// Creates a new instance of the Logger.
//
// Parameters:
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

// An alias for the NewLogger function.
func New(loggerInput *NewLoggerInput) *Logger {
	return NewLogger(loggerInput)
}

// Converts a given log level from an int to a string or from a number as a string to a log level
// name. For example: "0" -> "DEBUG" or "DEBUG" -> "0".
//
// If the log level is a string then it is used as is as the  map key. If the log level is an int
// then it is converted to a string and used as the map key. If none of the above return a value
// then the LogLevel property of the Logger is used as the map key.
func (l *Logger) ConvertLogLevel(level interface{}) string {
	switch v := level.(type) {
	// If the log level is a string then use it as is as the map key.
	case string:
		return logLevels[v]
	// If the log level is an int then convert it to a string and use it as the map key.
	case int:
		return logLevels[fmt.Sprintf("%d", v)]
	}

	// If none of the above return a value then try the LogLevel property of the Logger as a map key.
	return logLevels[fmt.Sprintf("%d", l.LogLevel)]
}

func GetCallingFunctionName() string {
	pc, _, _, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()

	return path.Base(funcName)
}

// #region Debug

// Outputs a message to the console using the DEBUG log level, fmt.Println, and coloring the
// output bright blue.
func (l *Logger) Debug(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	if l.LogLevel == 0 {
		logPrefix := l.getLogPrefix(&l.LogLevel, ctx)
		fmt.Println(aurora.BrightBlue(logPrefix + msg))
	}
}

// Outputs a message to the console using the DEBUG log level, fmt.Printf, and coloring the
// output bright blue.
func (l *Logger) Debugf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightBlue(msg))
}

// #endregion Debug

// #region Info

// Outputs a message to the console using the INFO log level, fmt.Println, and coloring the
// output cyan.
func (l *Logger) Info(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	if l.LogLevel <= 1 {
		logLevel := 1
		logPrefix := l.getLogPrefix(&logLevel, ctx)
		fmt.Println(aurora.Cyan(logPrefix + "" + msg))
	}
}

// Outputs a message to the console using the INFO log level, fmt.Printf, and coloring the
// output cyan.
func (l *Logger) Infof(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.Cyan(msg))
}

// #endregion Info

// #region Warn

// Outputs a message to the console using the WARN log level, fmt.Println, and coloring the
// output bright yellow.
func (l *Logger) Warn(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	if l.LogLevel <= 2 {
		logLevel := 2
		logPrefix := l.getLogPrefix(&logLevel, ctx)
		fmt.Println(aurora.BrightYellow(logPrefix + msg))
	}
}

// Outputs a message to the console using the WARN log level, fmt.Printf, and coloring the
// output bright yellow.
func (l *Logger) Warnf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightYellow(msg))
}

// #endregion Warn

// #region Error

// Outputs a message to the console using the ERROR log level, fmt.Println, and coloring the
// output bright red.
func (l *Logger) Error(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	if l.LogLevel <= 3 {
		logLevel := 3
		logPrefix := l.getLogPrefix(&logLevel, ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

// Outputs a message to the console using the ERROR log level, fmt.Printf, and coloring the
// output bright red.
func (l *Logger) Errorf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightRed(msg))
}

// #endregion Error

// #region Fatal

// Outputs a message to the console using the FATAL log level, fmt.Println, and coloring the
// output bright red.
func (l *Logger) Fatal(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	if l.LogLevel <= 4 {
		logLevel := 4
		logPrefix := l.getLogPrefix(&logLevel, ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

// Outputs a message to the console using the FATAL log level, fmt.Printf, and coloring the
// output bright red.
func (l *Logger) Fatalf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightRed(msg))
}

// #endregion Fatal

// #region Success

// Outputs a message to the console using fmt.Println, and coloring the output bright green.
func (l *Logger) Success(msg string, ctx *string) {
	if ctx == nil {
		funcName := GetCallingFunctionName()
		ctx = &funcName
	}

	logPrefix := "[" + l.AppName + "-SUCCESS#" + l.getCtx(ctx) + "]: "
	fmt.Println(aurora.BrightGreen(logPrefix + msg))
}

// Outputs a message to the console using the SUCCESS log level, fmt.Printf, and coloring the
// output bright green.
func (l *Logger) Successf(format string, input ...interface{}) {
	msg := fmt.Sprintf(format, input...)

	fmt.Printf(format, aurora.BrightGreen(msg))
}

// #endregion Success

// #endregion Exported
