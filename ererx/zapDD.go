package ererx

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

func (z *ZapDD) elog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(z.skip).Debug(msg, fields...) //这里当出错时还是 DEBUG 级别的
}

func (z *ZapDD) dlog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(z.skip).Debug(msg, fields...)
}
