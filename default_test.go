package logger

import (
	"bytes"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	tests := []struct {
		want, fmt     string
		a, b, c       interface{}
		edebug, debug bool
	}{
		{" INFO 1 2 3\n", "%d %d %d", 1, 2, 3, false, false},
		{" DEBUG 1 2 3\n", "%d %d %d", 1, 2, 3, true, true},
		{"", "%d %d %d", 1, 2, 3, false, true},
	}
	for _, tc := range tests {
		t.Run(tc.want, func(t *testing.T) {
			buf := &bytes.Buffer{}
			l := New(WithWriter(buf), WithTimeFormat(""))
			Set(l)
			EnableDebug(tc.edebug)
			if tc.debug {
				Debug(tc.fmt, tc.a, tc.b, tc.c)
			} else {
				Info(tc.fmt, tc.a, tc.b, tc.c)
			}
			if got := buf.String(); got != tc.want {
				t.Fatalf("expected %q; got %q", tc.want, got)
			}
		})
	}
}

func TestNilLogger(t *testing.T) {
	l := nilLogger{}
	EnableDebug(true)
	l.Debug("no message: %s", "abc")
	l.Info("no message: %s", "abc")
}
