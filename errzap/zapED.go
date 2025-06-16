package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type zapED struct{}

func NewLogED() errlog.Log {
	return &zapED{}
}

func (z *zapED) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(msg, fields...)
}

func (z *zapED) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
