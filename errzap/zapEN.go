package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// zapEN implements Log interface with ERROR level for errors and no logging for debug
// Logs errors as ERROR but suppresses debug level output
// Useful when you want to log errors but not Is/Ise match results
//
// zapEN 实现 Log 接口，错误使用 ERROR 级别，调试不记录日志
// 将错误记录为 ERROR 但抑制调试级别输出
// 用于需要记录错误但不记录 Is/Ise 匹配结果的场景
type zapEN struct{}

// NewLogEN creates a new zapEN log instance
// Returns Log interface with ERROR-only strategy
//
// NewLogEN 创建新的 zapEN 日志实例
// 返回带有仅 ERROR 策略的 Log 接口
func NewLogEN() errlog.Log {
	return &zapEN{}
}

// ErrorLog logs error level messages using zaplog
// Auto skips 2 stack frames for accurate log location
//
// ErrorLog 使用 zaplog 记录错误级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (z *zapEN) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(msg, fields...)
}

// DebugLog suppresses debug level logging
// Does nothing to avoid verbose output
//
// DebugLog 抑制调试级别日志
// 不执行任何操作以避免详细输出
func (z *zapEN) DebugLog(msg string, fields ...zap.Field) {
	// DO NOTHING: NOT PRINT LOG IN THIS CASE
	// 不执行任何操作：此情况不记录日志
}
