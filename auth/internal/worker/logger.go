package worker

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Print(level zerolog.Level, args ...any) {
	log.WithLevel(level).Msg(fmt.Sprint(args...))
}

// Debug logs a message at Debug level.
func (l *Logger) Debug(args ...any) {
	l.Print(zerolog.DebugLevel, args...)
}

// Info logs a message at Info level.
func (l *Logger) Info(args ...any) {
	l.Print(zerolog.InfoLevel, args...)
}

// Warn logs a message at Warning level.
func (l *Logger) Warn(args ...any) {
	l.Print(zerolog.WarnLevel, args...)
}

// Error logs a message at Error level.
func (l *Logger) Error(args ...any) {
	l.Print(zerolog.ErrorLevel, args...)
}

// Fatal logs a message at Fatal level
// and process will exit with status set to 1.
func (l *Logger) Fatal(args ...any) {
	l.Print(zerolog.FatalLevel, args...)
}
