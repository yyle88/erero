package erero

import (
	"github.com/yyle88/erero/ererx"
)

// 这个变量专用于以下的函数，而以下的函数能够方便使用
// 主要是以前先提供的函数，接下来才封的变量，就保留了这些函数，以保持兼容性，而且确实是直接使用函数能省掉麻烦
var eee = ererx.NewErero(ererx.NewZapED(3))

func Ero(erx error) error {
	return eee.Ero(erx)
}

func New(message string) error {
	return eee.New(message)
}

func WithMessage(err error, message string) error {
	return eee.WithMessage(err, message)
}

func Wms(err error, message string) error {
	return eee.Wms(err, message)
}

func Wme(err error, message string) error {
	return eee.Wme(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return eee.WithMessagef(err, format, args...)
}

func Wmf(err error, format string, args ...interface{}) error {
	return eee.Wmf(err, format, args...)
}

func Errorf(format string, args ...interface{}) error {
	return eee.Errorf(format, args...)
}

func Erf(format string, args ...interface{}) error {
	return eee.Erf(format, args...)
}

// Is 就是判定错误是否和另一个错误相等的，把这个函数也写过来，能避免引用 errors 包避免包名冲突还要使用别名
func Is(erx, target error) bool {
	return eee.Is(erx, target)
}

// Ise 就是 Is 的扩展，特别是当 Ise(err, gorm.ErrRecordNotFound) 的时候，假如确实是找不到就打印DEBUG日志，而其它错误就打印ERROR日志
func Ise(erx, target error) bool {
	return eee.Ise(erx, target)
}

func As(erx error, target any) bool {
	return eee.As(erx, target)
}

func Ase(erx error, target any) bool {
	return eee.Ase(erx, target)
}

// Wro 在很多时候虽然有错误，但是懒得写 WithMessage 的时候就可以直接用这个函数，避免代码里都是 WithMessage(err, "wrong") 稍微简化代码
func Wro(erx error) error {
	return eee.Wro(erx)
}

// Pan 就是当有错时 panic，相当于断言，这样也能避免在外面再判断有没有出错，择机使用吧，有些正式项目里还是不能随便panic的
func Pan(erx error) {
	eee.Pan(erx)
}
