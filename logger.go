package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	_ Logger = nilLogger{}
	_ Logger = &Impl{}
)

const (
	// DefaultTimeFormat used by the logger.
	// It is set to ISO-8601.
	DefaultTimeFormat = time.RFC3339
)

// Logger defines the main interface for logging.
type Logger interface {
	Info(string, ...interface{})  // issue a info message
	Debug(string, ...interface{}) // issue a debugging message
}

type nilLogger struct{}

func (nilLogger) Info(string, ...interface{})  {}
func (nilLogger) Debug(string, ...interface{}) {}

// Impl is a basic implementation of the Logger interface.
type Impl struct {
	out  io.Writer
	tfmt string
}

// Info writes a info message.
func (l *Impl) Info(format string, args ...interface{}) {
	l.write("INFO", format, args...)
}

// Debug writes a debug message if not disabled.
func (l *Impl) Debug(format string, args ...interface{}) {
	l.write("DEBUG", format, args...)
}

func (l *Impl) write(kind, format string, args ...interface{}) {
	fmt.Fprintf(l.out, "[%s] %s %s\n",
		kind,
		time.Now().Format(l.tfmt),
		fmt.Sprintf(format, args...))
}

// WithTimeFormat sets the time format of the logger.
func WithTimeFormat(tfmt string) func(*Impl) {
	return func(l *Impl) {
		l.tfmt = tfmt
	}
}

// WithWriter sets the writer of the logger.
func WithWriter(out io.Writer) func(*Impl) {
	return func(l *Impl) {
		l.out = out
	}
}

// New returns a new Logger interface.
// It uses stderr, has no prefix and uses the DefaultTimeFormat.
func New(args ...func(*Impl)) *Impl {
	l := &Impl{os.Stderr, DefaultTimeFormat}
	for _, arg := range args {
		arg(l)
	}
	return l
}
