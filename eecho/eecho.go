package eecho

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Log interface {
	ErrorLog(msg string, fields ...zap.Field)
	DebugLog(msg string, fields ...zap.Field)
}

type ErrorsEcho struct {
	log Log
}

func NewErrorsEcho(log Log) *ErrorsEcho {
	return &ErrorsEcho{log: log}
}

// Ero logs the provided error and returns it.
// Ero 记录提供的错误并返回该错误。
func (x *ErrorsEcho) Ero(erx error) error {
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

// New creates a new error with the given message, logs it, and returns the error.
// New 使用给定的消息创建一个新的错误，记录该错误并返回。
func (x *ErrorsEcho) New(message string) error {
	erx := errors.New(message)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

// WithMessage adds a custom message to the existing error, logs it, and returns the new error.
// WithMessage 给现有的错误添加自定义消息，记录该错误并返回。
func (x *ErrorsEcho) WithMessage(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

// WithMessagef adds a formatted message to the existing error, logs it, and returns the new error.
// WithMessagef 给现有的错误添加格式化消息，记录该错误并返回。
func (x *ErrorsEcho) WithMessagef(err error, format string, args ...interface{}) error {
	erx := errors.WithMessagef(err, format, args...)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

// Errorf creates a new formatted error, logs it, and returns the error.
// Errorf 创建一个新的格式化错误，记录该错误并返回。
func (x *ErrorsEcho) Errorf(format string, args ...interface{}) error {
	erx := errors.Errorf(format, args...)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

// Is checks if the provided error is equal to the target error.
// Is 检查提供的错误是否与目标错误相等。
func (x *ErrorsEcho) Is(erx, target error) bool {
	return errors.Is(erx, target)
}

// Ise checks if the provided error is equal to the target error, logging DEBUG for matching errors and ERROR for others.
// Ise 检查提供的错误是否与目标错误相等，对于匹配的错误记录 DEBUG 日志，其他错误记录 ERROR 日志。
func (x *ErrorsEcho) Ise(erx, target error) bool {
	if errors.Is(erx, target) {
		if erx != nil {
			x.log.DebugLog("DEBUG", zap.Error(erx))
		}
		return true
	}
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return false
}

// As checks if the provided error can be cast to the target type.
// As 检查提供的错误是否可以转换为目标类型。
func (x *ErrorsEcho) As(erx error, target any) bool {
	return errors.As(erx, target)
}

// Ase checks if the provided error can be cast to the target type, logging DEBUG for successful casts and ERROR for others.
// Ase 检查提供的错误是否可以转换为目标类型，成功转换时记录 DEBUG 日志，其他情况记录 ERROR 日志。
func (x *ErrorsEcho) Ase(erx error, target any) bool {
	if errors.As(erx, target) {
		if erx != nil {
			x.log.DebugLog("DEBUG", zap.Error(erx), zap.Any("target", target))
		}
		return true
	}
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return false
}

// Wro adds a default "wrong" message to the provided error, logs it, and returns the new error.
// Wro 为提供的错误添加默认的 "wrong" 消息，记录该错误并返回。
func (x *ErrorsEcho) Wro(err error) error {
	erx := errors.WithMessage(err, "wrong")
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *ErrorsEcho) Wrap(err error, message string) error {
	erx := errors.Wrap(err, message)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *ErrorsEcho) Wrapf(err error, format string, args ...interface{}) error {
	erx := errors.Wrapf(err, format, args...)
	if erx != nil {
		x.log.ErrorLog("ERROR", zap.Error(erx))
	}
	return erx
}
