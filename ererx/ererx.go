package ererx

import (
	"github.com/pkg/errors"
	"github.com/yyle88/erero/erebz"
	"go.uber.org/zap"
)

type xlog interface {
	elog(msg string, fields ...zap.Field)
	dlog(msg string, fields ...zap.Field)
}

type Erero struct {
	xlog xlog
}

func NewErero(xlog xlog) *Erero {
	return &Erero{
		xlog: xlog,
	}
}

func (x *Erero) Ero(erx error) error {
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) New(message string) error {
	erx := errors.New(message)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) WithMessage(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) Wms(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) Wme(err error, message string) error {
	erx := errors.WithMessage(err, message)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) WithMessagef(err error, format string, args ...interface{}) error {
	erx := errors.WithMessagef(err, format, args...)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) Wmf(err error, format string, args ...interface{}) error {
	erx := errors.WithMessagef(err, format, args...)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) Errorf(format string, args ...interface{}) error {
	erx := errors.Errorf(format, args...)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) Erf(format string, args ...interface{}) error {
	erx := errors.Errorf(format, args...)
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

// Is 就是判定错误是否和另一个错误相等的，把这个函数也写过来，能避免引用 errors 包避免包名冲突还要使用别名
func (x *Erero) Is(erx, target error) bool {
	return errors.Is(erx, target)
}

// Ise 就是 Is 的扩展，特别是当 Ise(err, gorm.ErrRecordNotFound) 的时候，假如确实是找不到就打印DEBUG日志，而其它错误就打印ERROR日志
func (x *Erero) Ise(erx, target error) bool {
	if errors.Is(erx, target) {
		if erx != nil { //当命中错误的时候就只打印 DEBUG 日志
			x.xlog.dlog("DEBUG", zap.Error(erx))
		}
		return true
	}
	if erx != nil { //而当不命中的时候就打印 ERROR 日志
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return false
}

func (x *Erero) As(erx error, target any) bool {
	return errors.As(erx, target)
}

func (x *Erero) Ase(erx error, target any) bool {
	if errors.As(erx, target) {
		if erx != nil { //当命中错误的时候就只打印 DEBUG 日志
			x.xlog.dlog("DEBUG", zap.Error(erx), zap.Any("target", target))
		}
		return true
	}
	if erx != nil { //而当不命中的时候就打印 ERROR 日志
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return false
}

// Wro 在很多时候虽然有错误，但是懒得写 WithMessage 的时候就可以直接用这个函数，避免代码里都是 WithMessage(err, "wrong") 稍微简化代码
func (x *Erero) Wro(err error) error {
	erx := errors.WithMessage(err, "wrong")
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
	}
	return erx
}

func (x *Erero) WroWebRequest(erx error) error {
	wre := erebz.WithMessage(erebz.NewWebRequestError(), erx, "wrong")
	if wre != nil {
		x.xlog.elog("ERROR", zap.Error(wre))
	}
	return wre
}

func (x *Erero) WroHttpStatus(code int, erx error) error {
	sub := erebz.NewHttpStatusError(code)
	hse := erebz.WithMessage(sub, erx, "wrong")
	if hse != nil {
		x.xlog.elog("ERROR", zap.Error(hse))
	}
	return hse
}

func (x *Erero) WroNotExist(erx error) error {
	nee := erebz.WithMessage(erebz.NewNotExistError(), erx, "wrong")
	if nee != nil {
		x.xlog.elog("ERROR", zap.Error(nee))
	}
	return nee
}

func (x *Erero) WroDatabase(erx error) error {
	dbe := erebz.WithMessage(erebz.NewDatabaseError(), erx, "wrong")
	if dbe != nil {
		x.xlog.elog("ERROR", zap.Error(dbe))
	}
	return dbe
}

func (x *Erero) WroRedis(erx error) error {
	rde := erebz.WithMessage(erebz.NewRedisError(), erx, "wrong")
	if rde != nil {
		x.xlog.elog("ERROR", zap.Error(rde))
	}
	return rde
}

func (x *Erero) WroCommonBiz(reason string, erx error) error {
	cme := erebz.WithMessage(erebz.NewCommonBizError(reason), erx, "wrong")
	if cme != nil {
		x.xlog.elog("ERROR", zap.Error(cme))
	}
	return cme
}

func (x *Erero) Pan(erx error) {
	if erx != nil {
		x.xlog.elog("ERROR", zap.Error(erx))
		panic(erx)
	}
}
