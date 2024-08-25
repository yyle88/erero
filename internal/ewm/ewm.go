// Package ewm
// COPY FROM github.com/pkg/errors@v0.9.1/errors.go type withMessage
package ewm

import (
	"fmt"
	"io"
)

func WithMessage(err error, message string) *Ewm {
	if err == nil {
		return nil
	}
	return &Ewm{
		cause: err,
		msg:   message,
	}
}

func WithMessagef(err error, format string, args ...interface{}) *Ewm {
	if err == nil {
		return nil
	}
	return &Ewm{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
}

type Ewm struct {
	cause error
	msg   string
}

func (w *Ewm) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *Ewm) Cause() error  { return w.cause }

func (w *Ewm) Unwrap() error { return w.cause }

func (w *Ewm) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", w.Cause())
			_, _ = io.WriteString(s, w.msg)
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, w.Error())
	}
}
