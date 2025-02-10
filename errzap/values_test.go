package errzap

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/errlog"
)

func TestErrors_Implement(t *testing.T) {
	var _ errlog.Log = NewZapDD(2)
	var _ errlog.Log = NewZapED(2)
	var _ errlog.Log = NewZapEN(2)
	var _ errlog.Log = NewZapWD(2)
}

func TestErrors_Ero(t *testing.T) {
	require.Error(t, ED.Ero(errors.New("abc")))
	require.Error(t, WD.Ero(errors.New("abc")))
	require.Error(t, DD.Ero(errors.New("abc")))
	require.Error(t, EN.Ero(errors.New("abc")))
}

func TestErrors_Ise(t *testing.T) {
	err := errors.New("abc")

	require.True(t, ED.Ise(err, err))
	require.True(t, WD.Ise(err, err))
	require.True(t, DD.Ise(err, err))
	t.Log("--")
	require.True(t, EN.Ise(err, err)) // NOT PRINT LOG
	t.Log("--")

	require.False(t, ED.Ise(err, errors.New("abc")))
	require.False(t, WD.Ise(err, errors.New("abc")))
	require.False(t, DD.Ise(err, errors.New("abc")))
	require.False(t, EN.Ise(err, errors.New("abc")))
}
