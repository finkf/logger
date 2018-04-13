// Package logger tries to define a
// simplistic logging interface.
package logger

import (
	"io"
	"log"
)

// Logger defines the main interface for logging.
type Logger interface {
	Printf(string, ...interface{}) // issue a debugging message
	Debugf(string, ...interface{}) // issue a info message
	Debug(bool)                    // enable/disable debug output
}

// Nil returns a logger instance with no output
func Nil() Logger {
	return nilLogger{}
}

type nilLogger struct{}

func (nilLogger) Printf(string, ...interface{}) {}
func (nilLogger) Debugf(string, ...interface{}) {}
func (nilLogger) Debug(bool)                    {}

// New returns a new Logger interface that uses go's log.Logger.
func New(out io.Writer, prefix string, flag int) Logger {
	return &defaultLogger{l: log.New(out, prefix, flag)}
}

type defaultLogger struct {
	l     *log.Logger
	debug bool
}

func (l *defaultLogger) Printf(fmt string, args ...interface{}) {
	l.l.Printf("[INFO] "+fmt, args...)
}

func (l *defaultLogger) Debugf(fmt string, args ...interface{}) {
	if !l.debug {
		return
	}
	l.l.Printf("[DEBG] "+fmt, args...)
}

func (l *defaultLogger) Debug(debug bool) {
	l.debug = debug
}
