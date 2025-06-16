package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type zapWD struct{}

func NewLogWD() errlog.Log {
	return &zapWD{}
}

func (z *zapWD) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Warn(msg, fields...) //这里出错时打 WARING 级日志
}

func (z *zapWD) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Debug(msg, fields...)
}
