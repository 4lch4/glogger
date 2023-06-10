package glogger

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

type Logger struct {
	// 0 = Debug, 1 = Info, 2 = Warn, 3 = Error, 4 = Fatal
	LogLevel int
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(msg string) {
	if l.LogLevel == 0 {
		fmt.Println(aurora.BrightBlack(msg))
	}
}

func (l *Logger) Info(msg string) {
	if l.LogLevel <= 1 {
		fmt.Println(aurora.Yellow(msg))
	}
}

func (l *Logger) Warn(msg string) {
	if l.LogLevel <= 2 {
		fmt.Println(aurora.BrightYellow(msg))
	}
}

func (l *Logger) Error(msg string) {
	if l.LogLevel <= 3 {
		fmt.Println(aurora.BrightRed(msg))
	}
}

func (l *Logger) Fatal(msg string) {
	if l.LogLevel <= 4 {
		fmt.Println(aurora.BrightRed(msg))
	}
}

func (l *Logger) Success(msg string) {
	fmt.Println(aurora.BrightGreen(msg))
}
