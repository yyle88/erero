// Package errzap provides pre-configured Errlog instances with different log level strategies
// Offers flexible error logging with ERROR, WARNING, and DEBUG combinations
// Integrates with zaplog for structured logging with configurable severity levels
// Each instance uses different logging patterns for error and Is/Ise operations
//
// errzap 提供带有不同日志级别策略的预配置 Errlog 实例
// 提供带有 ERROR、WARNING 和 DEBUG 组合的灵活错误日志
// 与 zaplog 集成以支持可配置严重级别的结构化日志
// 每个实例对错误和 Is/Ise 操作使用不同的日志模式
package errzap

import (
	"github.com/yyle88/erero/errlog"
)

// ED logs ERROR level for errors and DEBUG level when errors match with Is check
// ED 当出错时记录 ERROR 级别日志，当错误满足 Is 检查时记录 DEBUG 级别日志
var ED = errlog.NewErrlog(NewLogED())

// EN logs ERROR level for errors and no logging when errors match with Is check
// EN 当出错时记录 ERROR 级别日志，当错误满足 Is 检查时不记录日志
var EN = errlog.NewErrlog(NewLogEN())

// WD logs WARNING level for errors and DEBUG level when errors match with Is check
// WD 当出错时记录 WARNING 级别日志，当错误满足 Is 检查时记录 DEBUG 级别日志
var WD = errlog.NewErrlog(NewLogWD())

// DD logs DEBUG level for both errors and Is check matches
// DD 对错误和 Is 检查匹配都记录 DEBUG 级别日志
var DD = errlog.NewErrlog(NewLogDD())
