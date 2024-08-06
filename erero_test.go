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

func TestIse(t *testing.T) {
	{
		var erx = errors.New("wrong")
		//当错误匹配时，就打印调试日志
		require.True(t, Ise(erx, erx))
	}
	{
		//当错误不匹配，就打印错误日志
		require.False(t, Ise(errors.New("wrong"), errors.New("wrong")))
	}
}

func TestWro(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, WithMessage(erx, "wrong"))
	require.Error(t, Wro(erx)) //和前一行等效，能稍微省点代码
}
