package erero

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEEcho(t *testing.T) {
	require.Error(t, runExample())
}

func runExample() error {
	return op.New("abc")
}
