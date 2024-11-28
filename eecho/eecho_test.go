package eecho

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type echoExample struct {
	prefix string
}

func newEchoExample(prefix string) *echoExample {
	return &echoExample{prefix: prefix}
}

func (e *echoExample) ErrorLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(2).Error(e.prefix+msg, fields...)
}

func (e *echoExample) DebugLog(msg string, fields ...zap.Field) {
	zaplog.LOGS.Pn(2).Debug(e.prefix+msg, fields...)
}

var caseEEcho *EEcho

func TestMain(m *testing.M) {
	caseEEcho = NewEEcho(newEchoExample("echo-message:"))
	m.Run()
}

func TestEro(t *testing.T) {
	require.Error(t, caseEEcho.Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, caseEEcho.New("wrong"))
}

func TestIse(t *testing.T) {
	var erx = errors.New("wrong")

	//当错误匹配时，就打印调试日志
	require.True(t, caseEEcho.Ise(erx, erx))
	//当错误不匹配，就打印错误日志
	require.False(t, caseEEcho.Ise(erx, errors.New("wrong")))
}

func TestWro(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseEEcho.WithMessage(erx, "wrong"))
	require.Error(t, caseEEcho.Wro(erx)) //和前一行等效，能稍微省点代码
}
