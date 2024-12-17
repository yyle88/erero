package erzap

import (
	"github.com/yyle88/erero/eecho"
)

// ED 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var ED = eecho.NewErrorsEcho(NewZapED(2))

// EN 当出错时打印 ERROR 级别的日志，当错误满足 Is 的时候会不打印日志
var EN = eecho.NewErrorsEcho(NewZapEN(2))

// WD 当出错时打印 WARNING 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var WD = eecho.NewErrorsEcho(NewZapWD(2))

// DD 当出错时打印 DEBUG 级别的日志，当错误满足 Is 的时候会打印 DEBUG 级别日志
var DD = eecho.NewErrorsEcho(NewZapDD(2))
