package errzap

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

func (z *ZapEN) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(z.skip).Error(msg, fields...)
}

func (z *ZapEN) DebugLog(msg string, fields ...zap.Field) {
	//DO NOTHING: NOT PRINT LOG IN THIS CASE
}
