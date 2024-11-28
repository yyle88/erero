package erzap

import (
	"github.com/yyle88/erero/eecho"
)

// ED 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var ED = eecho.NewEEcho(NewZapED(2))

// EN 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会不打印日志
var EN = eecho.NewEEcho(NewZapEN(2))

// WD 当出错时打印 WARNING 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var WD = eecho.NewEEcho(NewZapWD(2))

// DD 当出错时打印 DEBUG 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var DD = eecho.NewEEcho(NewZapDD(2))
