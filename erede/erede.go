package erede

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Ero(erx error) error {
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}

func New(message string) error {
	erx := errors.New(message)
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}

func WithMessage(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}

func WithMessagef(err error, format string, args ...interface{}) error {
	erx := errors.WithMessagef(err, format, args...)
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}

func Errorf(format string, args ...interface{}) error {
	erx := errors.Errorf(format, args...)
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}

// Is 就是判定错误是否和另一个错误相等的，把这个函数也写过来，能避免引用 errors 包避免包名冲突还要使用别名
func Is(erx, target error) bool {
	return errors.Is(erx, target)
}

// Ise 就是 Is 的扩展，特别是当 Ise(err, gorm.ErrRecordNotFound) 的时候，假如确实是找不到就打印DEBUG日志，而其它错误就打印ERROR日志
func Ise(erx, target error) bool {
	if errors.Is(erx, target) {
		if erx != nil { //当命中错误的时候就只打印 DEBUG 日志
			zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
		}
		return true
	}
	if erx != nil { //而当不命中的时候就打印 DEBUG 日志
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return false
}

// Wro 在很多时候虽然有错误，但是懒得写 WithMessage 的时候就可以直接用这个函数，避免代码里都是 WithMessage(err, "wrong") 稍微简化代码
func Wro(err error) error {
	erx := errors.WithMessage(err, "wrong")
	if erx != nil {
		zaplog.LOGS.P1.Debug("DEBUG", zap.Error(erx))
	}
	return erx
}
