package errzap_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/errzap"
)

func TestEro(t *testing.T) {
	require.Error(t, errzap.ED.Ero(errors.New("abc")))
	require.Error(t, errzap.WD.Ero(errors.New("abc")))
	require.Error(t, errzap.DD.Ero(errors.New("abc")))
	require.Error(t, errzap.EN.Ero(errors.New("abc")))
}

func TestIse(t *testing.T) {
	err := errors.New("abc")

	require.True(t, errzap.ED.Ise(err, err))
	require.True(t, errzap.WD.Ise(err, err))
	require.True(t, errzap.DD.Ise(err, err))
	t.Log("--")
	require.True(t, errzap.EN.Ise(err, err)) // NOT PRINT LOG
	t.Log("--")

	require.False(t, errzap.ED.Ise(err, errors.New("abc")))
	require.False(t, errzap.WD.Ise(err, errors.New("abc")))
	require.False(t, errzap.DD.Ise(err, errors.New("abc")))
	require.False(t, errzap.EN.Ise(err, errors.New("abc")))
}
