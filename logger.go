package glogger

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

type Logger struct {
	// 0 = Debug / 1 = Info / 2 = Warn / 3 = Error / 4 = Fatal / 5 = Panic
	LogLevel int

	// The name of the application using the logger.
	AppName string
}

func getCtx(ctx string) string {
	if ctx == "" {
		return "Glogger"
	} else {
		return ctx
	}
}

// Creates a new instance of the Logger struct. Parameters:
//
// - appName: The name of the application using the logger. If nil, the default value is "Glogger".
//
// - logLevel: 0 = Debug / 1 = Info / 2 = Warn / 3 = Error / 4 = Fatal / 5 = Panic. If nil, the default value is 0.
func NewLogger(appName *string, logLevel *int) *Logger {
	defaultAppName := "Glogger"
	defaultLogLevel := 0

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
	case 5:
		return "PANIC"
	}

	return "UNKNOWN"
}

func (l *Logger) getLogPrefix(ctx string) string {
	return "[" + l.AppName + "-" + l.getLogLevel() + "]" + "[" + getCtx(ctx) + "]: "
}

func (l *Logger) Debug(msg string, ctx string) {
	if l.LogLevel == 0 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightBlue(logPrefix + msg))
	}
}

func (l *Logger) Info(msg string, ctx string) {
	if l.LogLevel <= 1 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.Cyan(logPrefix + "" + msg))
	}
}

func (l *Logger) Warn(msg string, ctx string) {
	if l.LogLevel <= 2 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightYellow(logPrefix + msg))
	}
}

func (l *Logger) Error(msg string, ctx string) {
	if l.LogLevel <= 3 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

func (l *Logger) Fatal(msg string, ctx string) {
	if l.LogLevel <= 4 {
		logPrefix := l.getLogPrefix(ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

func (l *Logger) Success(msg string, ctx string) {
	logPrefix := l.getLogPrefix(ctx)
	fmt.Println(aurora.BrightGreen(logPrefix + msg))
}
