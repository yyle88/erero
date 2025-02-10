package erero

import (
	"github.com/yyle88/erero/errlog"
	"github.com/yyle88/erero/errzap"
)

var op = errlog.NewErrlog(errzap.NewZapED(3))
