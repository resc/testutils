package testutils

import (
	"github.com/ReSc/fmt"
	"runtime"
	"strings"
	"testing"
)

type T struct {
	*testing.T
	counter int
	Fact    interface{}
}

func NewT(t *testing.T) T {
	return T{t, 0, nil}
}

func (t *T) Equal(expected, actual interface{}) {
	t.assert(expected == actual, "Expected %+v, got %+v", expected, actual)
}

func (t *T) Assert(b bool, format string, msg ...interface{}) {
	t.assert(b, format, msg...)
}

func (t *T) assert(b bool, format string, msg ...interface{}) {
	t.counter++
	if !b {
		t.Errorf(t.decorate(fmt.String(format, msg...)))
	}
}

func (t *T) decorate(s string) string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}

	return fmt.String("%s at \n%s:%d: assert #%d fact: %+v", s, file, line, t.counter, t.Fact)
}
