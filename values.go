package erero

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// ero is the default Errlog instance with zapLog implementation
// Auto configured with ERROR and DEBUG level logging
//
// ero 是带有 zapLog 实现的默认 Errlog 实例
// 自动配置 ERROR 和 DEBUG 级别日志
var ero = errlog.NewErrlog(newLog())

// SetLog replaces the default log implementation with a custom one
// Enables runtime log backend switching
//
// SetLog 用自定义实现替换默认日志实现
// 启用运行时日志后端切换
func SetLog(log errlog.Log) {
	ero.SetLog(log)
}

// zapLog implements Log interface using zaplog as the backend
// Provides structured logging with ERROR and DEBUG levels
// Auto skips 3 stack frames for accurate log location
//
// zapLog 使用 zaplog 作为后端实现 Log 接口
// 提供带有 ERROR 和 DEBUG 级别的结构化日志
// 自动跳过 3 个堆栈帧以获得准确的日志位置
type zapLog struct{}

// newLog creates a new zapLog instance
// Returns default Log implementation for erero package
//
// newLog 创建新的 zapLog 实例
// 返回 erero 包的默认 Log 实现
func newLog() *zapLog {
	return &zapLog{}
}

// ErrorLog logs error level messages using zaplog
// Auto skips 3 stack frames for accurate log location
//
// ErrorLog 使用 zaplog 记录错误级别消息
// 自动跳过 3 个堆栈帧以获得准确的日志位置
func (z *zapLog) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(3).Error(msg, fields...)
}

// DebugLog logs debug level messages using zaplog
// Auto skips 3 stack frames for accurate log location
//
// DebugLog 使用 zaplog 记录调试级别消息
// 自动跳过 3 个堆栈帧以获得准确的日志位置
func (z *zapLog) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(3).Debug(msg, fields...)
}
