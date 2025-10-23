package errzap_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero/errzap"
)

// TestEro verifies err logging across these pre-configured strategies
// Tests ED (ERROR-DEBUG), WD (WARNING-DEBUG), DD (DEBUG-DEBUG), EN (ERROR-None)
//
// TestEro 验证所有预配置日志级别策略的错误记录
// 测试 ED（ERROR-DEBUG）、WD（WARNING-DEBUG）、DD（DEBUG-DEBUG）、EN（ERROR-None）
func TestEro(t *testing.T) {
	require.Error(t, errzap.ED.Ero(errors.New("abc")))
	require.Error(t, errzap.WD.Ero(errors.New("abc")))
	require.Error(t, errzap.DD.Ero(errors.New("abc")))
	require.Error(t, errzap.EN.Ero(errors.New("abc")))
}

// TestIse verifies condition-based logging across these strategies
// Tests both matching (logs DEBUG) and non-matching (logs ERROR) cases
//
// TestIse 验证所有策略的条件日志行为
// 测试匹配（记录 DEBUG）和不匹配（记录 ERROR）的情况
func TestIse(t *testing.T) {
	err := errors.New("abc")

	// When errors match, all strategies log DEBUG except EN (no log)
	// 当错误匹配时，除 EN（无日志）外所有策略都记录 DEBUG
	require.True(t, errzap.ED.Ise(err, err))
	require.True(t, errzap.WD.Ise(err, err))
	require.True(t, errzap.DD.Ise(err, err))
	t.Log("--")

	// EN strategy does not log when errors match
	// EN 策略在错误匹配时不记录日志
	require.True(t, errzap.EN.Ise(err, err))
	t.Log("--")

	// When errors do not match, all strategies log their configured error level
	// 当错误不匹配时，所有策略都记录其配置的错误级别
	require.False(t, errzap.ED.Ise(err, errors.New("abc")))
	require.False(t, errzap.WD.Ise(err, errors.New("abc")))
	require.False(t, errzap.DD.Ise(err, errors.New("abc")))
	require.False(t, errzap.EN.Ise(err, errors.New("abc")))
}
