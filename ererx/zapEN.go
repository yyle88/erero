package ererx

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ZapEN struct {
	skip int
}

func NewZapEN(skip int) *ZapEN {
	return &ZapEN{skip: skip}
}

func (z *ZapEN) elog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(z.skip).Error(msg, fields...)
}

func (z *ZapEN) dlog(msg string, fields ...zap.Field) {
	//DO NOTHING: NOT PRINT LOG IN THIS CASE
}
