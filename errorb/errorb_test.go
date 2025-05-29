package errorb

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewErrorOrNotExist(t *testing.T) {
	erb := NewErrorOrNotExist(errors.New("wrong"))
	require.Error(t, erb.Cause)
	require.False(t, erb.NotExist)
}

func TestNewErrorIsNotExist(t *testing.T) {
	erb := NewErrorIsNotExist()
	require.Error(t, erb.Cause)
	require.True(t, erb.NotExist)
}

func TestNewErrorOrSuccess(t *testing.T) {
	erb := NewErrorOrSuccess(errors.New("wrong"))
	require.Error(t, erb.Cause)
	require.False(t, erb.Success)
}

func TestNewErrorIsSuccess(t *testing.T) {
	erb := NewErrorIsSuccess()
	require.Error(t, erb.Cause)
	require.True(t, erb.Success)
}

func TestNewErrorOrIgnore(t *testing.T) {
	erb := NewErrorOrIgnore(errors.New("wrong"))
	require.Error(t, erb.Cause)
	require.False(t, erb.Ignore)
}

func TestNewErrorIsIgnore(t *testing.T) {
	erb := NewErrorIsIgnore()
	require.Error(t, erb.Cause)
	require.True(t, erb.Ignore)
}

func TestNewErrorOrRetryable(t *testing.T) {
	erb := NewErrorOrRetryable(errors.New("wrong"))
	require.Error(t, erb.Cause)
	require.False(t, erb.Retryable)
}

func TestNewErrorIsRetryable(t *testing.T) {
	erb := NewErrorIsRetryable()
	require.Error(t, erb.Cause)
	require.True(t, erb.Retryable)
}

func TestNewErrorOrRecoverable(t *testing.T) {
	erb := NewErrorOrRecoverable(errors.New("wrong"))
	require.Error(t, erb.Cause)
	require.False(t, erb.Recoverable)
}

func TestNewErrorIsRecoverable(t *testing.T) {
	erb := NewErrorIsRecoverable()
	require.Error(t, erb.Cause)
	require.True(t, erb.Recoverable)
}
