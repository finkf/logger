package logger

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	tests := []struct {
		want, fmt     string
		a, b, c       interface{}
		edebug, debug bool
	}{
		{"[INFO] 1 2 3\n", "%d %d %d", 1, 2, 3, false, false},
		{"[DEBG] 1 2 3\n", "%d %d %d", 1, 2, 3, true, true},
		{"", "%d %d %d", 1, 2, 3, false, true},
	}
	for _, tc := range tests {
		t.Run(tc.want, func(t *testing.T) {
			buf := &bytes.Buffer{}
			l := New(buf, "", 0)
			Set(l)
			if !reflect.DeepEqual(l, Get()) {
				t.Fatalf("invalid default logger")
			}
			Debug(tc.edebug)
			if tc.debug {
				Debugf(tc.fmt, tc.a, tc.b, tc.c)
			} else {
				Printf(tc.fmt, tc.a, tc.b, tc.c)
			}
			if got := buf.String(); got != tc.want {
				t.Fatalf("expected %q; got %q", tc.want, got)
			}
		})
	}
}

func TestNilLogger(t *testing.T) {
	l := Nil()
	l.Debug(true)
	l.Debugf("no message: %s", "abc")
	l.Printf("no message: %s", "abc")
}
