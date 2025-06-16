package errlog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type zapLog struct {
	prefix string
}

func newLog(prefix string) *zapLog {
	return &zapLog{prefix: prefix}
}

func (e *zapLog) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(e.prefix+":"+msg, fields...)
}

func (e *zapLog) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(e.prefix+":"+msg, fields...)
}

func TestSetLog(t *testing.T) {
	var ero = errlog.NewErrlog(newLog("prefix-v1x"))
	require.Error(t, ero.New("wrong"))

	ero.SetLog(newLog("prefix-v2x"))
	require.Error(t, ero.New("wrong"))
}
