package ewm

import (
	"testing"

	"github.com/pkg/errors"
)

func TestWithMessage(t *testing.T) {
	t.Log(errors.New("wrong"))
	t.Log(WithMessage(errors.New("wrong"), "wrong"))
}

func TestWithMessagef(t *testing.T) {
	t.Log(errors.New("wrong"))
	t.Log(WithMessagef(errors.New("wrong"), "erx=%s", "wrong"))
}
