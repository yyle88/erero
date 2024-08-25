package erebz

import (
	"github.com/yyle88/erero/internal/ewm"
)

type WebRequestError struct {
	*supEwmBase
}

func NewWebRequestError() *WebRequestError {
	return &WebRequestError{
		supEwmBase: &supEwmBase{},
	}
}

type HttpStatusError struct {
	*supEwmBase
	statusCode int //就是http的状态码
}

func NewHttpStatusError(statusCode int) *HttpStatusError {
	return &HttpStatusError{
		supEwmBase: &supEwmBase{},
		statusCode: statusCode,
	}
}

func (a *HttpStatusError) StatusCode() int {
	return a.statusCode
}

type DatabaseError struct {
	*supEwmBase
}

func NewDatabaseError() *DatabaseError {
	return &DatabaseError{
		supEwmBase: &supEwmBase{},
	}
}

type RedisError struct {
	*supEwmBase
}

func NewRedisError() *RedisError {
	return &RedisError{
		supEwmBase: &supEwmBase{},
	}
}

type CommonBizError struct {
	*supEwmBase
	reason string
}

func NewCommonBizError(reason string) *CommonBizError {
	return &CommonBizError{
		supEwmBase: &supEwmBase{},
		reason:     reason,
	}
}

func (e *CommonBizError) Reason() string {
	return e.reason
}

type NotExistError struct {
	*supEwmBase
}

func NewNotExistError() *NotExistError {
	return &NotExistError{
		supEwmBase: &supEwmBase{},
	}
}

type supEwmBase struct {
	*ewm.Ewm
}

func (a *supEwmBase) setEwm(sub *ewm.Ewm) {
	a.Ewm = sub
}

type supInterface interface {
	error
	setEwm(sub *ewm.Ewm)
}

func WithMessage[T supInterface](sup T, err error, message string) error {
	sub := ewm.WithMessage(err, message)
	if sub == nil {
		return nil
	}
	sup.setEwm(sub)
	return sup
}

func WithMessagef[T supInterface](sup T, err error, format string, args ...interface{}) error {
	sub := ewm.WithMessagef(err, format, args...)
	if sub == nil {
		return nil
	}
	sup.setEwm(sub)
	return sup
}
