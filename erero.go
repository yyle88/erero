// Package erero provides err handling with auto context and location logging
// Wraps errors with stack traces and logs them using configurable settings
// Avoids naming conflicts with Go standard errors package and famous err handling packages
// Integrates seamlessly with zap logging engine providing structured output
//
// erero 提供带有自动上下文和位置记录的错误处理
// 用堆栈跟踪包装错误并使用可配置的日志级别记录它们
// 避免与 Go 标准 errors 包和流行错误处理包的命名冲突
// 与 zap 日志记录器无缝集成以进行结构化错误记录
package erero

// Ero logs the provided err and returns it
// Auto records err with stack trace context
//
// Ero 记录提供的错误并返回它
// 自动记录带有堆栈跟踪上下文的错误
func Ero(err error) error {
	return ero.Ero(err)
}

// New creates a new err with the given message and logs it
// Auto records err with stack trace context
//
// New 使用给定消息创建新错误并记录它
// 自动记录带有堆栈跟踪上下文的错误
func New(message string) error {
	return ero.New(message)
}

// WithMessage annotates the provided err with a custom message and logs it
// Auto records err with enhanced context
//
// WithMessage 用自定义消息注释提供的错误并记录它
// 自动记录带有增强上下文的错误
func WithMessage(err error, message string) error {
	return ero.WithMessage(err, message)
}

// WithMessagef annotates the provided err with a formatted message and logs it
// Auto records err with formatted context
//
// WithMessagef 用格式化消息注释提供的错误并记录它
// 自动记录带有格式化上下文的错误
func WithMessagef(err error, format string, args ...interface{}) error {
	return ero.WithMessagef(err, format, args...)
}

// Errorf creates a new formatted err and logs it
// Auto records err with stack trace context
//
// Errorf 创建新的格式化错误并记录它
// 自动记录带有堆栈跟踪上下文的错误
func Errorf(format string, args ...interface{}) error {
	return ero.Errorf(format, args...)
}

// Is checks if the provided err matches the target
// Standard err comparison without logging
//
// Is 检查提供的错误是否与目标错误匹配
// 标准错误比较不记录日志
func Is(err, target error) bool {
	return ero.Is(err, target)
}

// Ise checks if the provided err matches the target with condition-based logging
// Logs DEBUG when matching and ERROR otherwise
//
// Ise 检查提供的错误是否与目标错误匹配并条件记录日志
// 匹配时记录 DEBUG，否则记录 ERROR
func Ise(err, target error) bool {
	return ero.Ise(err, target)
}

// As checks if the provided err can be cast to the target type
// Standard err type assertion without logging
//
// As 检查提供的错误是否可转换为目标类型
// 标准错误类型断言不记录日志
func As(err error, target any) bool {
	return ero.As(err, target)
}

// Ase checks if the provided err can be cast to the target type with condition-based logging
// Logs DEBUG when casting succeeds and ERROR otherwise
//
// Ase 检查提供的错误是否可转换为目标类型并条件记录日志
// 成功转换时记录 DEBUG，否则记录 ERROR
func Ase(err error, target any) bool {
	return ero.Ase(err, target)
}

// Wro annotates the provided err with a default "wrong" message and logs it
// Quick shorthand to WithMessage with predefined message
//
// Wro 用默认的 "wrong" 消息注释提供的错误并记录它
// 使用预定义消息的 WithMessage 便捷简写
func Wro(err error) error {
	return ero.Wro(err)
}

// Wrap wraps the provided err with a message and logs it
// Includes complete stack trace in wrapped err
//
// Wrap 用消息包装提供的错误并记录它
// 在包装的错误中包含完整堆栈跟踪
func Wrap(err error, message string) error {
	return ero.Wrap(err, message)
}

// Wrapf wraps the provided err with a formatted message and logs it
// Includes complete stack trace in wrapped err
//
// Wrapf 用格式化消息包装提供的错误并记录它
// 在包装的错误中包含完整堆栈跟踪
func Wrapf(err error, format string, args ...interface{}) error {
	return ero.Wrapf(err, format, args...)
}

// Join combines multiple errors into one and logs it
// Uses Go standard errors.Join with logging
//
// Join 将多个错误合并为单个错误并记录它
// 使用 Go 标准 errors.Join 并记录日志
func Join(errs ...error) error {
	return ero.Join(errs...)
}

// Joins combines a slice of errors into one and logs it
// Quick wrapper to Join accepting err slices
//
// Joins 将错误切片合并为单个错误并记录它
// 接受错误切片的 Join 便捷包装
func Joins(errs []error) error {
	return ero.Joins(errs)
}
