package ereru

import (
	"github.com/yyle88/erero/ererx"
)

var ere3x = ererx.NewErerx(ererx.NewZapED(3))

func As[T any](erx error) (*T, bool) {
	var target *T
	return target, ere3x.As(erx, &target)
}

func Ase[T any](erx error) (*T, bool) {
	var target *T
	return target, ere3x.Ase(erx, &target)
}

func WroWebRequest(erx error) error {
	return ere3x.WroWebRequest(erx)
}

func WroHttpStatus(code int, erx error) error {
	return ere3x.WroHttpStatus(code, erx)
}

func WroNotExist(erx error) error {
	return ere3x.WroNotExist(erx)
}

func WroDatabase(erx error) error {
	return ere3x.WroDatabase(erx)
}

func WroRedis(erx error) error {
	return ere3x.WroRedis(erx)
}

func WroCommonBiz(reason string, erx error) error {
	return ere3x.WroCommonBiz(reason, erx)
}
