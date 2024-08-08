package erero

import "github.com/yyle88/erero/ererx"

// ED 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var ED = ererx.NewErerx(ererx.NewZapED(2))

// EN 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会不打印日志
var EN = ererx.NewErerx(ererx.NewZapEN(2))

// WD 当出错时打印 WARNING 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var WD = ererx.NewErerx(ererx.NewZapWD(2))

// DD 当出错时打印 DEBUG 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var DD = ererx.NewErerx(ererx.NewZapDD(2))
