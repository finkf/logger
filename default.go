package logger

import (
	"log"
	"os"
)

var def Logger

func init() {
	def = New(os.Stderr, "", log.LstdFlags)
}

// Default sets the default logger
func Default(l Logger) {
	def = l
}

// Printf issues an info-level message with the default logger.
func Printf(fmt string, args ...interface{}) {
	def.Printf(fmt, args...)
}

// Debugf issues a debug-level message with the default logger.
func Debugf(fmt string, args ...interface{}) {
	def.Debugf(fmt, args...)
}

// Debug enables or disables the debug output for the default logger.
func Debug(debug bool) {
	def.Debug(debug)
}
