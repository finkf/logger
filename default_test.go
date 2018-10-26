package logger

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestDefaultLogger(t *testing.T) {
	tformat := "20060102"
	tests := []struct {
		want, fmt     string
		a, b, c       interface{}
		edebug, debug bool
	}{
		{"%s INFO 1 2 3\n", "%d %d %d", 1, 2, 3, false, false},
		{"%s DEBUG 1 2 3\n", "%d %d %d", 1, 2, 3, true, true},
		{"", "%d %d %d", 1, 2, 3, false, true},
	}
	for _, tc := range tests {
		t.Run(tc.want, func(t *testing.T) {
			buf := &bytes.Buffer{}
			l := New(WithWriter(buf), WithTimeFormat(tformat))
			Set(l)
			EnableDebug(tc.edebug)
			if tc.debug {
				Debug(tc.fmt, tc.a, tc.b, tc.c)
			} else {
				Info(tc.fmt, tc.a, tc.b, tc.c)
			}
			if tc.want != "" {
				tc.want = fmt.Sprintf(tc.want, time.Now().Format(tformat))
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
