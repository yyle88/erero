package ererx

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var ee *Erero

func TestMain(m *testing.M) {
	ee = NewErero(NewZapED(2))
	m.Run()

	ee = NewErero(NewZapWD(2))
	m.Run()

	ee = NewErero(NewZapDD(2))
	m.Run()
}

func TestEro(t *testing.T) {
	require.Error(t, ee.Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, ee.New("wrong"))
}

func TestWms(t *testing.T) {
	require.Error(t, ee.Wms(ee.New("wrong"), "msg"))
}

func TestIse(t *testing.T) {
	{
		var erx = errors.New("wrong")
		//当错误匹配时，就打印调试日志
		require.True(t, ee.Ise(erx, erx))
	}
	{
		//当错误不匹配，就打印错误日志
		require.False(t, ee.Ise(errors.New("wrong"), errors.New("wrong")))
	}
}

func TestWro(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, ee.WithMessage(erx, "wrong"))
	require.Error(t, ee.Wro(erx)) //和前一行等效，能稍微省点代码
}

func TestErerx_ErxBiz(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, ee.WithMessage(erx, "wrong"))
	{
		wbe := ee.WroWebRequest(erx)
		t.Log(wbe)
		require.Error(t, wbe) //和前一行等效，能稍微省点代码
	}
	{
		ehc := ee.WroHttpStatus(http.StatusBadRequest, erx)
		t.Log(ehc)
		require.Error(t, ehc) //和前一行等效，能稍微省点代码
	}
	{
		dbe := ee.WroDatabase(erx)
		t.Log(dbe)
		require.Error(t, dbe) //和前一行等效，能稍微省点代码
	}
	{
		rbe := ee.WroRedis(erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}

	{
		rbe := ee.WroCommonBiz("wa", erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}

	{
		rbe := ee.WroNotExist(erx)
		t.Log(rbe)
		require.Error(t, rbe) //和前一行等效，能稍微省点代码
	}
}
