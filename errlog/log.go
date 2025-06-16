package errlog

import "go.uber.org/zap"

type Log interface {
	ErrorLog(msg string, fields ...zap.Field)
	DebugLog(msg string, fields ...zap.Field)
}

// SetLog assigns a logger to the Errlog instance.
// SetLog 给 Errlog 实例分配日志记录器。
func (x *Errlog) SetLog(log Log) {
	x.log = log
}
