package erero

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/erero/errzap"
)

var ero = errlog.NewErrlog(errzap.NewZapED(3))

func SetLog(log errlog.Log) {
	ero = errlog.NewErrlog(log)
}
