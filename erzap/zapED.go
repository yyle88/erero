package erzap

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ZapED struct {
	skip int
}

func NewZapED(skip int) *ZapED {
	return &ZapED{skip: skip}
}

func (z *ZapED) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(z.skip).Error(msg, fields...)
}

func (z *ZapED) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(z.skip).Debug(msg, fields...)
}
