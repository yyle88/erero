package erero

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Ero(erx error) error {
	if erx != nil {
		zaplog.LOGS.P1.Error("ERROR", zap.Error(erx))
	}
	return erx
}

func New(message string) error {
	erx := errors.New(message)
	if erx != nil {
		zaplog.LOGS.P1.Error("ERROR", zap.Error(erx))
	}
	return erx
}

func WithMessage(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		zaplog.LOGS.P1.Error("ERROR", zap.Error(erx))
	}
	return erx
}

func WithMessagef(err error, format string, args ...interface{}) error {
	erx := errors.WithMessagef(err, format, args...)
	if erx != nil {
		zaplog.LOGS.P1.Error("ERROR", zap.Error(erx))
	}
	return erx
}

func Errorf(format string, args ...interface{}) error {
	erx := errors.Errorf(format, args...)
	if erx != nil {
		zaplog.LOGS.P1.Error("ERROR", zap.Error(erx))
	}
	return erx
}

// Is 就是判定错误是否和另一个错误相等的，把这个函数也写过来，能避免引用 errors 包避免包名冲突还要使用别名
func Is(err, target error) bool {
	return errors.Is(err, target)
}
