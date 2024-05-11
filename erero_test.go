package erero

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestEro(t *testing.T) {
	require.Error(t, Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, New("wrong"))
}

func TestWithMessage(t *testing.T) {
	require.Error(t, WithMessage(New("wrong"), "msg"))
}
