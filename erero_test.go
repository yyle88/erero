package erero_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func TestEro(t *testing.T) {
	require.Error(t, erero.Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, erero.New("wrong"))
}

func TestIse(t *testing.T) {
	var err = errors.New("wrong")

	//当错误匹配时，就打印调试日志
	require.True(t, erero.Ise(err, err))
	//当错误不匹配，就打印错误日志
	require.False(t, erero.Ise(err, errors.New("wrong")))
}

func TestWro(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessage(err, "wrong"))
	require.Error(t, erero.Wro(err)) //和前一行等效，能稍微省点代码
}

func TestWithMessage(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessage(err, "wrong"))
}

func TestWithMessagef(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessagef(err, "wrong reason=%s", "reason"))
}

func TestWrap(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.Wrap(err, "wrong"))
}

func TestWrapf(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.Wrapf(err, "wrong reason=%s", "reason"))
}
