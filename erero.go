package erero

import (
	"github.com/yyle88/erero/ererx"
)

// 这个变量专用于以下的函数，而以下的函数能够方便使用
// 主要是以前先提供的函数，接下来才封的变量，就保留了这些函数，以保持兼容性，而且确实是直接使用函数能省掉麻烦
var e3z = ererx.NewErerx(ererx.NewZapED(3))

func Ero(erx error) error {
	return e3z.Ero(erx)
}

func New(message string) error {
	return e3z.New(message)
}

func WithMessage(err error, message string) error {
	return e3z.WithMessage(err, message)
}

func Wms(err error, message string) error {
	return e3z.Wms(err, message)
}

func Wme(err error, message string) error {
	return e3z.Wme(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return e3z.WithMessagef(err, format, args...)
}

func Wmf(err error, format string, args ...interface{}) error {
	return e3z.Wmf(err, format, args...)
}

func Errorf(format string, args ...interface{}) error {
	return e3z.Errorf(format, args...)
}

func Erf(format string, args ...interface{}) error {
	return e3z.Erf(format, args...)
}

// Is 就是判定错误是否和另一个错误相等的，把这个函数也写过来，能避免引用 errors 包避免包名冲突还要使用别名
func Is(erx, target error) bool {
	return e3z.Is(erx, target)
}

// Ise 就是 Is 的扩展，特别是当 Ise(err, gorm.ErrRecordNotFound) 的时候，假如确实是找不到就打印DEBUG日志，而其它错误就打印ERROR日志
func Ise(erx, target error) bool {
	return e3z.Ise(erx, target)
}

func As(erx error, target any) bool {
	return e3z.As(erx, target)
}

func Ase(erx error, target any) bool {
	return e3z.Ase(erx, target)
}

// Wro 在很多时候虽然有错误，但是懒得写 WithMessage 的时候就可以直接用这个函数，避免代码里都是 WithMessage(err, "wrong") 稍微简化代码
func Wro(err error) error {
	return e3z.Wro(err)
}
