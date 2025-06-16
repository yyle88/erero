package erero

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var ero = errlog.NewErrlog(newLog())

func SetLog(log errlog.Log) {
	ero.SetLog(log)
}

type zapLog struct{}

func newLog() *zapLog {
	return &zapLog{}
}

func (z *zapLog) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(3).Error(msg, fields...)
}

func (z *zapLog) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(3).Debug(msg, fields...)
}
