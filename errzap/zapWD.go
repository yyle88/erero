package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// zapWD implements Log interface with WARNING level for errors and DEBUG level for debug
// Logs errors as WARNING and debug messages as DEBUG
// Useful when errors are not critical but should still be tracked
//
// zapWD 实现 Log 接口，错误使用 WARNING 级别，调试使用 DEBUG 级别
// 将错误记录为 WARNING，调试消息记录为 DEBUG
// 用于错误不严重但仍应跟踪的场景
type zapWD struct{}

// NewLogWD creates a new zapWD log instance
// Returns Log interface with WARNING-DEBUG strategy
//
// NewLogWD 创建新的 zapWD 日志实例
// 返回带有 WARNING-DEBUG 策略的 Log 接口
func NewLogWD() errlog.Log {
	return &zapWD{}
}

// ErrorLog logs warning level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// ErrorLog 使用 zaplog 记录警告级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapWD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Warn(msg, fields...)
}

// DebugLog logs debug level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// DebugLog 使用 zaplog 记录调试级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapWD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
