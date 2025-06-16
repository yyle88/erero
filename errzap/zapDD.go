package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type zapDD struct{}

func NewLogDD() errlog.Log {
	return &zapDD{}
}

func (z *zapDD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...) //这里当出错时还是 DEBUG 级别的
}

func (z *zapDD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
