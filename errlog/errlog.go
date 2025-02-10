package errlog

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Log interface {
	ErrorLog(msg string, fields ...zap.Field)
	DebugLog(msg string, fields ...zap.Field)
}

type Errlog struct {
	log Log
}

func NewErrlog(log Log) *Errlog {
	return &Errlog{log: log}
}

// Ero logs the provided error and returns it.
// Ero 记录提供的错误并返回该错误。
func (x *Errlog) Ero(err error) error {
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// New creates a new error with the given message, logs it, and returns the error.
// New 使用给定的消息创建一个新的错误，记录该错误并返回。
func (x *Errlog) New(message string) error {
	err := errors.New(message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessage adds a custom message to the existing error, logs it, and returns the new error.
// WithMessage 给现有的错误添加自定义消息，记录该错误并返回。
func (x *Errlog) WithMessage(err error, message string) error {
	err = errors.WithMessage(err, message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessagef adds a formatted message to the existing error, logs it, and returns the new error.
// WithMessagef 给现有的错误添加格式化消息，记录该错误并返回。
func (x *Errlog) WithMessagef(err error, format string, args ...interface{}) error {
	err = errors.WithMessagef(err, format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Errorf creates a new formatted error, logs it, and returns the error.
// Errorf 创建一个新的格式化错误，记录该错误并返回。
func (x *Errlog) Errorf(format string, args ...interface{}) error {
	err := errors.Errorf(format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Is checks if the provided error is equal to the target error.
// Is 检查提供的错误是否与目标错误相等。
func (x *Errlog) Is(err, target error) bool {
	return errors.Is(err, target)
}

// Ise checks if the provided error is equal to the target error, logging DEBUG for matching errors and ERROR for others.
// Ise 检查提供的错误是否与目标错误相等，对于匹配的错误记录 DEBUG 日志，其他错误记录 ERROR 日志。
func (x *Errlog) Ise(err, target error) bool {
	if errors.Is(err, target) {
		if err != nil {
			x.log.DebugLog("DEBUG", zap.Error(err))
		}
		return true
	}
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return false
}

// As checks if the provided error can be cast to the target type.
// As 检查提供的错误是否可以转换为目标类型。
func (x *Errlog) As(err error, target any) bool {
	return errors.As(err, target)
}

// Ase checks if the provided error can be cast to the target type, logging DEBUG for successful casts and ERROR for others.
// Ase 检查提供的错误是否可以转换为目标类型，成功转换时记录 DEBUG 日志，其他情况记录 ERROR 日志。
func (x *Errlog) Ase(err error, target any) bool {
	if errors.As(err, target) {
		if err != nil {
			x.log.DebugLog("DEBUG", zap.Error(err), zap.Any("target", target))
		}
		return true
	}
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return false
}

// Wro adds a default "wrong" message to the provided error, logs it, and returns the new error.
// Wro 为提供的错误添加默认的 "wrong" 消息，记录该错误并返回。
func (x *Errlog) Wro(err error) error {
	err = errors.WithMessage(err, "wrong")
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

func (x *Errlog) Wrap(err error, message string) error {
	err = errors.Wrap(err, message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

func (x *Errlog) Wrapf(err error, format string, args ...interface{}) error {
	err = errors.Wrapf(err, format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}
