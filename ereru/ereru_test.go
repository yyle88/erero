package ereru

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/erebz"
)

func TestWro(t *testing.T) {
	erx := WroWebRequest(errors.New("wrong"))

	{
		wbe, ok := Ase[erebz.WebRequestError](erx)
		require.True(t, ok)
		t.Log(wbe)
		require.NotNil(t, wbe)
	}

	{
		dbe, ok := Ase[erebz.DatabaseError](erx)
		require.False(t, ok)
		t.Log(dbe)
		require.Nil(t, dbe)
	}
}
