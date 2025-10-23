package errlog

import (
	"errors"

	"go.uber.org/zap"
)

// Join combines multiple errors into a single error and logs it
// Uses Go standard errors.Join with auto logging
//
// Join 将多个错误合并为单个错误并记录它
// 使用 Go 标准 errors.Join 并自动记录日志
func (ero *Errlog) Join(errs ...error) error {
	err := errors.Join(errs...)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}

// Joins combines a slice of errors into a single error and logs it
// Convenient wrapper for Join accepting error slices
//
// Joins 将错误切片合并为单个错误并记录它
// 接受错误切片的 Join 便捷包装
func (ero *Errlog) Joins(errs []error) error {
	err := errors.Join(errs...)
	if err != nil {
		ero.log.ErrorLog("ERROR", zap.Error(err))
	}
	return err
}
