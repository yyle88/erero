package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// zapDD implements Log interface with DEBUG level for both errors and debug
// Logs both errors and debug messages as DEBUG
// Useful for development and testing when verbose output is needed
//
// zapDD 实现 Log 接口，错误和调试都使用 DEBUG 级别
// 将错误和调试消息都记录为 DEBUG
// 用于需要详细输出的开发和测试场景
type zapDD struct{}

// NewLogDD creates a new zapDD log instance
// Returns Log interface with DEBUG-DEBUG strategy
//
// NewLogDD 创建新的 zapDD 日志实例
// 返回带有 DEBUG-DEBUG 策略的 Log 接口
func NewLogDD() errlog.Log {
	return &zapDD{}
}

// ErrorLog logs debug level messages for errors using zaplog
// Auto skips 2 stack frames for accurate log location
//
// ErrorLog 使用 zaplog 为错误记录调试级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapDD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}

// DebugLog logs debug level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// DebugLog 使用 zaplog 记录调试级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapDD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
