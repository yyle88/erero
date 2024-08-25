package erebz

import (
	"testing"

	"github.com/pkg/errors"
)

func TestWithMessage(t *testing.T) {
	erx := errors.New("wrong")
	t.Log(erx)
	wbe := WithMessage(NewWebRequestError(), erx, "wrong")
	t.Log(wbe)
}

func TestWithMessagef(t *testing.T) {
	erx := errors.New("wrong")
	t.Log(erx)
	wbe := WithMessagef(NewDatabaseError(), erx, "erx=%s", "wrong")
	t.Log(wbe)
}
