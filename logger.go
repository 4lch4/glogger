package glogger

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

type Logger struct {
	// 0 = Debug / 1 = Info / 2 = Warn / 3 = Error / 4 = Fatal / 5 = Panic
	LogLevel int
}

func getCtx(ctx string) string {
	if ctx == "" {
		return "Glogger"
	} else {
		return ctx
	}
}

func getLogPrefix(logLevel string, ctx string) string {
	return "[" + logLevel + "]" + "[" + getCtx(ctx) + "]: "
}

func NewLogger(logLevel *int) *Logger {
	defaultLogLevel := 0

	if logLevel == nil {
		logLevel = &defaultLogLevel
	}

	fmt.Println("logLevel: ", *logLevel)

	return &Logger{
		LogLevel: *logLevel,
	}
}

func (l *Logger) Debug(msg string, ctx string) {
	if l.LogLevel == 0 {
		logPrefix := getLogPrefix("DEBUG", ctx)
		fmt.Println(aurora.BrightBlue(logPrefix + msg))
	}
}

func (l *Logger) Info(msg string, ctx string) {
	if l.LogLevel <= 1 {
		logPrefix := getLogPrefix("INFO", ctx)
		fmt.Println(aurora.Yellow(logPrefix + "" + msg))
	}
}

func (l *Logger) Warn(msg string, ctx string) {
	if l.LogLevel <= 2 {
		logPrefix := getLogPrefix("WARN", ctx)
		fmt.Println(aurora.BrightYellow(logPrefix + msg))
	}
}

func (l *Logger) Error(msg string, ctx string) {

	if l.LogLevel <= 3 {
		logPrefix := getLogPrefix("ERROR", ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

func (l *Logger) Fatal(msg string, ctx string) {
	if l.LogLevel <= 4 {
		logPrefix := getLogPrefix("FATAL", ctx)
		fmt.Println(aurora.BrightRed(logPrefix + msg))
	}
}

func (l *Logger) Success(msg string, ctx string) {
	logPrefix := getLogPrefix("SUCCESS", ctx)
	fmt.Println(aurora.BrightGreen(logPrefix + msg))
}
