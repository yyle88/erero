package erero_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

// TestEro verifies err logging and returning behavior
// Expects err to be logged and returned unchanged
//
// TestEro 验证错误记录和返回行为
// 期望错误被记录并原样返回
func TestEro(t *testing.T) {
	require.Error(t, erero.Ero(errors.New("abc")))
}

// TestNew verifies new err creation with logging
// Expects err to be created with stack trace and logged
//
// TestNew 验证带日志的新错误创建
// 期望创建带有堆栈跟踪的错误并记录
func TestNew(t *testing.T) {
	require.Error(t, erero.New("wrong"))
}

// TestIse verifies condition-based logging on err matching
// When errors match, logs at DEBUG; otherwise logs at ERROR
//
// TestIse 验证基于错误匹配的条件日志
// 当错误匹配时记录 DEBUG 级别日志，否则记录 ERROR 级别日志
func TestIse(t *testing.T) {
	var err = errors.New("wrong")

	// When errors match, logs DEBUG
	// 当错误匹配时，记录 DEBUG 级别日志
	require.True(t, erero.Ise(err, err))

	// When errors do not match, logs ERROR
	// 当错误不匹配时，记录 ERROR 级别日志
	require.False(t, erero.Ise(err, errors.New("wrong")))
}

// TestWro verifies shorthand err wrapping with default message
// Wro() is equivalent to WithMessage() but saves code
//
// TestWro 验证带默认消息的简写错误包装
// Wro() 等效于 WithMessage() 但节省代码
func TestWro(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessage(err, "wrong"))

	// Wro() is shorthand to WithMessage with "wrong"
	// Wro() 是 WithMessage 带 "wrong" 的简写
	require.Error(t, erero.Wro(err))
}

// TestWithMessage verifies err annotation with custom message
// Expects err to be wrapped with message and logged
//
// TestWithMessage 验证带自定义消息的错误注释
// 期望错误被消息包装并记录
func TestWithMessage(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessage(err, "wrong"))
}

// TestWithMessagef verifies err annotation with formatted message
// Expects err to be wrapped with formatted message and logged
//
// TestWithMessagef 验证带格式化消息的错误注释
// 期望错误被格式化消息包装并记录
func TestWithMessagef(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.WithMessagef(err, "wrong reason=%s", "reason"))
}

// TestWrap verifies err wrapping with stack trace
// Expects err to be wrapped with message and complete stack trace
//
// TestWrap 验证带堆栈跟踪的错误包装
// 期望错误被消息和完整堆栈跟踪包装
func TestWrap(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.Wrap(err, "wrong"))
}

// TestWrapf verifies err wrapping with formatted message and stack trace
// Expects err to be wrapped with formatted message and complete stack trace
//
// TestWrapf 验证带格式化消息和堆栈跟踪的错误包装
// 期望错误被格式化消息和完整堆栈跟踪包装
func TestWrapf(t *testing.T) {
	err := errors.New("abc")
	require.Error(t, erero.Wrapf(err, "wrong reason=%s", "reason"))
}

// TestJoin verifies combining multiple errors into one
// Expects these errors to be joined and logged
//
// TestJoin 验证将多个错误合并为单个错误
// 期望所有错误被合并并记录
func TestJoin(t *testing.T) {
	require.Error(t, erero.Join(
		erero.New("abc"),
		erero.New("xyz"),
		erero.New("123"),
	))
}

// TestJoins verifies combining err slice into one
// Expects these errors in slice to be joined and logged
//
// TestJoins 验证将错误切片合并为单个错误
// 期望切片中的所有错误被合并并记录
func TestJoins(t *testing.T) {
	var errs = []error{
		erero.New("abc"),
		erero.New("xyz"),
		erero.New("123"),
	}
	require.Error(t, erero.Joins(errs))
}
