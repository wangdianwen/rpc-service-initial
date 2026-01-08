package logging

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
}

type SimpleLogger struct {
	*log.Logger
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *SimpleLogger) Info(msg string, args ...interface{}) {
	l.Printf("[INFO] "+msg, args...)
}

func (l *SimpleLogger) Error(msg string, args ...interface{}) {
	l.Printf("[ERROR] "+msg, args...)
}

func (l *SimpleLogger) Debug(msg string, args ...interface{}) {
	l.Printf("[DEBUG] "+msg, args...)
}

func (l *SimpleLogger) Warn(msg string, args ...interface{}) {
	l.Printf("[WARN] "+msg, args...)
}
