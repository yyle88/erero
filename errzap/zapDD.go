package errzap

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ZapDD struct {
	skip int
}

func NewZapDD(skip int) *ZapDD {
	return &ZapDD{skip: skip}
}

func (z *ZapDD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(z.skip).Debug(msg, fields...) //这里当出错时还是 DEBUG 级别的
}

func (z *ZapDD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(z.skip).Debug(msg, fields...)
}
