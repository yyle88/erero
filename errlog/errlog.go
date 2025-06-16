package errlog

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Errlog struct {
	log Log
}

func NewErrlog(log Log) *Errlog {
	return &Errlog{log: log}
}

// Ero logs the provided error and returns it.
// Ero 当有错误时返回错误并记录日志。
func (x *Errlog) Ero(err error) error {
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// New creates a new error with the given message and logs it.
// New 使用给定的消息创建新错误并打印日志。
func (x *Errlog) New(message string) error {
	err := errors.New(message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessage annotates the provided error with a custom message and logs it.
// WithMessage 给提供的错误添加自定义消息并打印日志。
func (x *Errlog) WithMessage(err error, message string) error {
	err = errors.WithMessage(err, message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessagef annotates the provided error with a formatted message and logs it.
// WithMessagef 给提供的错误添加格式化消息并打印日志。
func (x *Errlog) WithMessagef(err error, format string, args ...interface{}) error {
	err = errors.WithMessagef(err, format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Errorf creates a new formatted error and logs it.
// Errorf 创建新的格式化错误并打印日志。
func (x *Errlog) Errorf(format string, args ...interface{}) error {
	err := errors.Errorf(format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Is checks if the provided error matches the target error.
// Is 检查提供的错误是否与目标错误匹配。
func (x *Errlog) Is(err, target error) bool {
	return errors.Is(err, target)
}

// Ise checks if the provided error matches the target error, logging DEBUG for matches and ERROR otherwise.
// Ise 检查提供的错误是否与目标错误匹配，匹配时打印 DEBUG 日志，否则打印 ERROR 日志。
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
// As 检查提供的错误是否可转换为目标类型。
func (x *Errlog) As(err error, target any) bool {
	return errors.As(err, target)
}

// Ase checks if the provided error can be cast to the target type, logging DEBUG for successful casts and ERROR otherwise.
// Ase 检查提供的错误是否可转换为目标类型，成功转换时打印 DEBUG 日志，否则打印 ERROR 日志。
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

// Wro annotates the provided error with a default "wrong" message and logs it.
// Wro 给提供的错误添加默认的 "wrong" 消息并打印日志。
func (x *Errlog) Wro(err error) error {
	err = errors.WithMessage(err, "wrong")
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Wrap wraps the provided error with a message and logs it.
// Wrap 给提供的错误添加消息并包装后打印日志。
func (x *Errlog) Wrap(err error, message string) error {
	err = errors.Wrap(err, message)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Wrapf wraps the provided error with a formatted message and logs it.
// Wrapf 给提供的错误添加格式化消息并包装后打印日志。
func (x *Errlog) Wrapf(err error, format string, args ...interface{}) error {
	err = errors.Wrapf(err, format, args...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}
