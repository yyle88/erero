package errzap

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type ZapWD struct {
	skip int
}

func NewZapWD(skip int) *ZapWD {
	return &ZapWD{skip: skip}
}

func (z *ZapWD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(z.skip).Warn(msg, fields...) //这里出错时打 WARING 级日志
}

func (z *ZapWD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(z.skip).Debug(msg, fields...)
}
