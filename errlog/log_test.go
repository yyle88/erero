package errlog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// zapLog implements errlog.Log interface with custom prefix
// Adds prefix to messages enabling quick identification
//
// zapLog 实现带有自定义前缀的 errlog.Log 接口
// 为所有日志消息添加前缀以便于识别
type zapLog struct {
	prefix string // Log message prefix // 日志消息前缀
}

// newLog creates a new zapLog instance with specified prefix
// Returns Log interface implementation with prefix support
//
// newLog 创建带有指定前缀的新 zapLog 实例
// 返回带有前缀支持的 Log 接口实现
func newLog(prefix string) *zapLog {
	return &zapLog{prefix: prefix}
}

// ErrorLog logs err messages with prefix
// Auto skips 2 stack frames to get accurate location
//
// ErrorLog 记录带有前缀的错误级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (e *zapLog) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(e.prefix+":"+msg, fields...)
}

// DebugLog logs debug messages with prefix
// Auto skips 2 stack frames to get accurate location
//
// DebugLog 记录带有前缀的调试级别消息
// 自动跳过 2 个堆栈帧以获得准确的日志位置
func (e *zapLog) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(e.prefix+":"+msg, fields...)
}

// TestSetLog verifies dynamic log backend switching
// Expects log prefix to change after SetLog call
//
// TestSetLog 验证动态日志后端切换
// 期望调用 SetLog 后日志前缀改变
func TestSetLog(t *testing.T) {
	// Create Errlog with first prefix
	// 创建带有第一个前缀的 Errlog
	var ero = errlog.NewErrlog(newLog("prefix-v1x"))
	require.Error(t, ero.New("wrong"))

	// Switch to different prefix
	// 切换到不同的前缀
	ero.SetLog(newLog("prefix-v2x"))
	require.Error(t, ero.New("wrong"))
}
