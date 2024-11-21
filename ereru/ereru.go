package ereru

import (
	"github.com/yyle88/erero/ererx"
)

var eee = ererx.NewErero(ererx.NewZapED(3))

func As[T any](erx error) (*T, bool) {
	var target *T
	return target, eee.As(erx, &target)
}

func Ase[T any](erx error) (*T, bool) {
	var target *T
	return target, eee.Ase(erx, &target)
}

func WroWebRequest(erx error) error {
	return eee.WroWebRequest(erx)
}

func WroHttpStatus(code int, erx error) error {
	return eee.WroHttpStatus(code, erx)
}

func WroNotExist(erx error) error {
	return eee.WroNotExist(erx)
}

func WroDatabase(erx error) error {
	return eee.WroDatabase(erx)
}

func WroRedis(erx error) error {
	return eee.WroRedis(erx)
}

func WroCommonBiz(reason string, erx error) error {
	return eee.WroCommonBiz(reason, erx)
}
