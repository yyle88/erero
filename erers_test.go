package erero

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestErers(t *testing.T) {
	require.Error(t, ED.Ero(errors.New("abc")))
	require.Error(t, WD.Ero(errors.New("abc")))
	require.Error(t, DD.Ero(errors.New("abc")))
	require.Error(t, EN.Ero(errors.New("abc")))
}

func TestErers_Ise(t *testing.T) {
	a := errors.New("abc")
	b := errors.New("abc")

	require.True(t, ED.Ise(a, a))
	require.True(t, WD.Ise(a, a))
	require.True(t, DD.Ise(a, a))
	t.Log("--")
	require.True(t, EN.Ise(a, a)) // NOT PRINT LOG
	t.Log("--")

	require.False(t, ED.Ise(a, b))
	require.False(t, WD.Ise(a, b))
	require.False(t, DD.Ise(a, b))
	require.False(t, EN.Ise(a, b))
}
