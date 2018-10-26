package logger

import (
	"sync"
)

var (
	def Logger
	mut sync.Mutex
	deb bool
)

func init() {
	def = New()
}

func withMutex(f func()) {
	mut.Lock()
	defer mut.Unlock()
	f()
}

// Set sets the default logger.
func Set(l Logger) {
	withMutex(func() {
		def = l
	})
}

// Disable disables the logger.
func Disable() {
	Set(nilLogger{})
}

// Info issues an info-level message with the default logger.
func Info(fmt string, args ...interface{}) {
	withMutex(func() {
		def.Info(fmt, args...)
	})
}

// Debug issues a debug-level message with the default logger.
func Debug(fmt string, args ...interface{}) {
	withMutex(func() {
		if !deb {
			return
		}
		def.Debug(fmt, args...)
	})
}

// EnableDebug enables or disables the debug output for the default logger.
func EnableDebug(enable bool) {
	withMutex(func() {
		deb = enable
	})
}
