package erzap

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/eecho"
)

func TestErrors_Imp(t *testing.T) {
	var _ eecho.Log = NewZapDD(2)
	var _ eecho.Log = NewZapED(2)
	var _ eecho.Log = NewZapEN(2)
	var _ eecho.Log = NewZapWD(2)
}

func TestErrors_Ero(t *testing.T) {
	require.Error(t, ED.Ero(errors.New("abc")))
	require.Error(t, WD.Ero(errors.New("abc")))
	require.Error(t, DD.Ero(errors.New("abc")))
	require.Error(t, EN.Ero(errors.New("abc")))
}

func TestErrors_Ise(t *testing.T) {
	erx := errors.New("abc")

	require.True(t, ED.Ise(erx, erx))
	require.True(t, WD.Ise(erx, erx))
	require.True(t, DD.Ise(erx, erx))
	t.Log("--")
	require.True(t, EN.Ise(erx, erx)) // NOT PRINT LOG
	t.Log("--")

	require.False(t, ED.Ise(erx, errors.New("abc")))
	require.False(t, WD.Ise(erx, errors.New("abc")))
	require.False(t, DD.Ise(erx, errors.New("abc")))
	require.False(t, EN.Ise(erx, errors.New("abc")))
}
