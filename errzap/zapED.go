package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// zapED implements Log interface with ERROR level for errors and DEBUG level for debug
// Logs errors as ERROR and debug messages as DEBUG
// Useful when you want full error logging with debug level for Is/Ise matches
//
// zapED 实现 Log 接口，错误使用 ERROR 级别，调试使用 DEBUG 级别
// 将错误记录为 ERROR，调试消息记录为 DEBUG
// 用于需要完整错误日志记录并为 Is/Ise 匹配使用调试级别的场景
type zapED struct{}

// NewLogED creates a new zapED log instance
// Returns Log interface with ERROR-DEBUG strategy
//
// NewLogED 创建新的 zapED 日志实例
// 返回带有 ERROR-DEBUG 策略的 Log 接口
func NewLogED() errlog.Log {
	return &zapED{}
}

// ErrorLog logs error level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// ErrorLog 使用 zaplog 记录错误级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapED) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(msg, fields...)
}

// DebugLog logs debug level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// DebugLog 使用 zaplog 记录调试级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapED) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
