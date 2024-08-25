package ererx

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var ere2x *Ererx

func TestMain(m *testing.M) {
	ere2x = NewErerx(NewZapED(2))
	m.Run()

	ere2x = NewErerx(NewZapWD(2))
	m.Run()

	ere2x = NewErerx(NewZapDD(2))
	m.Run()
}

func TestEro(t *testing.T) {
	require.Error(t, ere2x.Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, ere2x.New("wrong"))
}

func TestWithMessage(t *testing.T) {
	require.Error(t, ere2x.WithMessage(ere2x.New("wrong"), "msg"))
}

func TestIse(t *testing.T) {
	{
		var erx = errors.New("wrong")
		//当错误匹配时，就打印调试日志
		require.True(t, ere2x.Ise(erx, erx))
	}
	{
		//当错误不匹配，就打印错误日志
		require.False(t, ere2x.Ise(errors.New("wrong"), errors.New("wrong")))
	}
}

func TestWro(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, ere2x.WithMessage(erx, "wrong"))
	require.Error(t, ere2x.Wro(erx)) //和前一行等效，能稍微省点代码
}

func TestErerx_ErxBiz(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, ere2x.WithMessage(erx, "wrong"))
	{
		wbe := ere2x.WroWebRequest(erx)
		t.Log(wbe)
		require.Error(t, wbe) //和前一行等效，能稍微省点代码
	}
	{
		ehc := ere2x.WroHttpStatus(http.StatusBadRequest, erx)
		t.Log(ehc)
		require.Error(t, ehc) //和前一行等效，能稍微省点代码
	}
	{
		dbe := ere2x.WroDatabase(erx)
		t.Log(dbe)
		require.Error(t, dbe) //和前一行等效，能稍微省点代码
	}
	{
		rbe := ere2x.WroRedis(erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}

	{
		rbe := ere2x.WroCommonBiz("wa", erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}

	{
		rbe := ere2x.WroNotExist(erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}
}
