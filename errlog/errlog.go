// Package errlog provides core error logging functionality with configurable log interfaces
// Combines pkg/errors stack traces with structured logging using zap fields
// Supports custom log implementations through Log interface abstraction
// Enables context-aware error handling with DEBUG and ERROR level logging
//
// errlog 提供带有可配置日志接口的核心错误日志功能
// 将 pkg/errors 堆栈跟踪与使用 zap 字段的结构化日志结合
// 通过 Log 接口抽象支持自定义日志实现
// 启用带有 DEBUG 和 ERROR 级别日志的上下文感知错误处理
package errlog

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Errlog wraps a Log interface to provide error handling with auto logging
// Core error processor that combines error operations with structured logging
// All error operations auto trigger appropriate log level based on context
//
// Errlog 包装 Log 接口以提供带有自动日志记录的错误处理
// 核心错误处理器，将错误操作与结构化日志结合
// 所有错误操作根据上下文自动触发适当的日志级别
type Errlog struct {
	log Log // Underlying log interface implementation // 底层日志接口实现
}

// NewErrlog creates a new Errlog instance with the provided log implementation
// Accepts any Log interface implementation for flexible logging backends
//
// NewErrlog 使用提供的日志实现创建新的 Errlog 实例
// 接受任何 Log 接口实现以支持灵活的日志后端
func NewErrlog(log Log) *Errlog {
	return &Errlog{log: log}
}

// Ero logs the provided error and returns it
// Auto records error when error is not nil
//
// Ero 记录提供的错误并返回它
// 当错误不为 nil 时自动记录错误
func (ero *Errlog) Ero(err error) error {
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// New creates a new error with the given message and logs it
// Auto records error with stack trace context
//
// New 使用给定消息创建新错误并记录它
// 自动记录带有堆栈跟踪上下文的错误
func (ero *Errlog) New(message string) error {
	err := errors.New(message)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessage annotates the provided error with a custom message and logs it
// Auto records error with enhanced context
//
// WithMessage 用自定义消息注释提供的错误并记录它
// 自动记录带有增强上下文的错误
func (ero *Errlog) WithMessage(err error, message string) error {
	err = errors.WithMessage(err, message)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// WithMessagef annotates the provided error with a formatted message and logs it
// Auto records error with formatted context
//
// WithMessagef 用格式化消息注释提供的错误并记录它
// 自动记录带有格式化上下文的错误
func (ero *Errlog) WithMessagef(err error, format string, args ...interface{}) error {
	err = errors.WithMessagef(err, format, args...)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Errorf creates a new formatted error and logs it
// Auto records error with stack trace context
//
// Errorf 创建新的格式化错误并记录它
// 自动记录带有堆栈跟踪上下文的错误
func (ero *Errlog) Errorf(format string, args ...interface{}) error {
	err := errors.Errorf(format, args...)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Is checks if the provided error matches the target error
// Standard error comparison without logging
//
// Is 检查提供的错误是否与目标错误匹配
// 标准错误比较不记录日志
func (ero *Errlog) Is(err, target error) bool {
	return errors.Is(err, target)
}

// Ise checks if the provided error matches the target error with conditional logging
// Logs DEBUG for matches and ERROR otherwise
//
// Ise 检查提供的错误是否与目标错误匹配并条件记录日志
// 匹配时记录 DEBUG，否则记录 ERROR
func (ero *Errlog) Ise(err, target error) bool {
	if errors.Is(err, target) {
		if err != nil {
			ero.log.DebugLog("DEBUG", zap.Error(err))
		}
		return true
	}
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return false
}

// As checks if the provided error can be cast to the target type
// Standard error type assertion without logging
//
// As 检查提供的错误是否可转换为目标类型
// 标准错误类型断言不记录日志
func (ero *Errlog) As(err error, target any) bool {
	return errors.As(err, target)
}

// Ase checks if the provided error can be cast to the target type with conditional logging
// Logs DEBUG for successful casts and ERROR otherwise
//
// Ase 检查提供的错误是否可转换为目标类型并条件记录日志
// 成功转换时记录 DEBUG，否则记录 ERROR
func (ero *Errlog) Ase(err error, target any) bool {
	if errors.As(err, target) {
		if err != nil {
			ero.log.DebugLog("DEBUG", zap.Error(err), zap.Any("target", target))
		}
		return true
	}
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return false
}

// Wro annotates the provided error with a default "wrong" message and logs it
// Convenient shorthand for WithMessage with predefined message
//
// Wro 用默认的 "wrong" 消息注释提供的错误并记录它
// 使用预定义消息的 WithMessage 便捷简写
func (ero *Errlog) Wro(err error) error {
	err = errors.WithMessage(err, "wrong")
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Wrap wraps the provided error with a message and logs it
// Includes full stack trace in wrapped error
//
// Wrap 用消息包装提供的错误并记录它
// 在包装的错误中包含完整堆栈跟踪
func (ero *Errlog) Wrap(err error, message string) error {
	err = errors.Wrap(err, message)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Wrapf wraps the provided error with a formatted message and logs it
// Includes full stack trace in wrapped error
//
// Wrapf 用格式化消息包装提供的错误并记录它
// 在包装的错误中包含完整堆栈跟踪
func (ero *Errlog) Wrapf(err error, format string, args ...interface{}) error {
	err = errors.Wrapf(err, format, args...)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}
