package logger

import (
	"fmt"
	"strings"
)

var LogLevel = map[string]int{
	"CRIT":  0,
	"ERROR": 1,
	"WARN":  2,
	"INFO":  3,
	"DEBUG": 4,
}

type Logger struct {
	level int
}

func New(level string) *Logger {
	return &Logger{
		level: LogLevel[strings.ToUpper(level)],
	}
}

func (l Logger) formatLog(msg string, level string) string {
	return fmt.Sprintf("[%s] %s", level, msg)
}

func (l Logger) writeLog(msg string) {
	fmt.Println(msg)
}

func (l Logger) handleMessage(level string, msg string) {
	levelNum, levelNumExist := LogLevel[level]
	if !levelNumExist && levelNum > l.level {
		return
	}
	log := l.formatLog(msg, level)
	l.writeLog(log)
}

func (l Logger) Debug(msg string) {
	l.handleMessage("DEBUG", msg)
}

func (l Logger) Info(msg string) {
	l.handleMessage("INFO", msg)
}

func (l Logger) Warn(msg string) {
	l.handleMessage("WARN", msg)
}

func (l Logger) Error(msg string) {
	l.handleMessage("ERROR", msg)
}

func (l Logger) Critical(msg string) {
	l.handleMessage("CRIT", msg)
}
