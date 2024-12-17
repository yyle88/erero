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

var caseErrorsEcho *ErrorsEcho

func TestMain(m *testing.M) {
	caseErrorsEcho = NewErrorsEcho(newEchoExample("echo-message:"))
	m.Run()
}

func TestEro(t *testing.T) {
	require.Error(t, caseErrorsEcho.Ero(errors.New("abc")))
}

func TestNew(t *testing.T) {
	require.Error(t, caseErrorsEcho.New("wrong"))
}

func TestIse(t *testing.T) {
	var erx = errors.New("wrong")

	//当错误匹配时，就打印调试日志
	require.True(t, caseErrorsEcho.Ise(erx, erx))
	//当错误不匹配，就打印错误日志
	require.False(t, caseErrorsEcho.Ise(erx, errors.New("wrong")))
}

func TestWro(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseErrorsEcho.WithMessage(erx, "wrong"))
	require.Error(t, caseErrorsEcho.Wro(erx)) //和前一行等效，能稍微省点代码
}

func TestWithMessage(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseErrorsEcho.WithMessage(erx, "wrong"))
}

func TestWithMessagef(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseErrorsEcho.WithMessagef(erx, "wrong reason=%s", "reason"))
}

func TestWrap(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseErrorsEcho.Wrap(erx, "wrong"))
}

func TestWrapf(t *testing.T) {
	erx := errors.New("abc")
	require.Error(t, caseErrorsEcho.Wrapf(erx, "wrong reason=%s", "reason"))
}
