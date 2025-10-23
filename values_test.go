package erero

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestExample verifies basic err creation
// Expects err to be created and logged using default logging
//
// TestExample 验证基本错误创建功能
// 期望使用默认日志记录器创建并记录错误
func TestExample(t *testing.T) {
	require.Error(t, runExample())
}

// runExample demonstrates basic err creation with logging
// Returns err created via default erero instance
//
// runExample 演示带日志的基本错误创建
// 返回通过默认 erero 实例创建的错误
func runExample() error {
	return ero.New("abc")
}
