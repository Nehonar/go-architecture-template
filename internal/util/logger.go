package util

import (
	"log"
)

// Logger is a simple wrapper that can be expanded or replaced
// with a more advanced logging library in real projects.
type Logger struct {
	*log.Logger
}

// NewLogger creates a new standard logger.
// In real projects, this may include log levels or structured logging.
func NewLogger() *Logger {
	return &Logger{Logger: log.Default()}
}
