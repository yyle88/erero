package errlog

import (
	"errors"

	"go.uber.org/zap"
)

// Join combines multiple errors into a single error and logs it.
// Join 将多个错误合并为单个错误并打印日志。
func (x *Errlog) Join(errs ...error) error {
	err := errors.Join(errs...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Joins combines a slice of errors into a single error and logs it.
// Joins 将错误切片合并为单个错误并打印日志。
func (x *Errlog) Joins(errs []error) error {
	err := errors.Join(errs...)
	if err != nil {
		x.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}
