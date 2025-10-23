package errlog

import "go.uber.org/zap"

// Log defines the interface for structured logging with ERROR and DEBUG levels
// Implementations must support zap.Field for structured log context
// Enables custom log backends while maintaining consistent error handling
//
// Log 定义带有 ERROR 和 DEBUG 级别的结构化日志接口
// 实现必须支持 zap.Field 以支持结构化日志上下文
// 启用自定义日志后端同时保持一致的错误处理
type Log interface {
	ErrorLog(msg string, fields ...zap.Field) // Log error level messages // 记录错误级别消息
	DebugLog(msg string, fields ...zap.Field) // Log debug level messages // 记录调试级别消息
}

// SetLog assigns a logger to the Errlog instance
// Enables dynamic log backend switching at runtime
//
// SetLog 给 Errlog 实例分配日志记录器
// 启用运行时动态日志后端切换
func (ero *Errlog) SetLog(log Log) {
	ero.log = log
}
