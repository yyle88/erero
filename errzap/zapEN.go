package errzap

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type zapEN struct{}

func NewLogEN() errlog.Log {
	return &zapEN{}
}

func (z *zapEN) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Skip(2).Error(msg, fields...)
}

func (z *zapEN) DebugLog(msg string, fields ...zap.Field) {
	//DO NOTHING: NOT PRINT LOG IN THIS CASE
}
