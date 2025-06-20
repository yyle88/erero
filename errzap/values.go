package errzap

import (
	"github.com/yyle88/erero/errlog"
)

// ED 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var ED = errlog.NewErrlog(NewLogED())

// EN 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会不打印日志
var EN = errlog.NewErrlog(NewLogEN())

// WD 当出错时打印 WARNING 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var WD = errlog.NewErrlog(NewLogWD())

// DD 当出错时打印 DEBUG 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var DD = errlog.NewErrlog(NewLogDD())
