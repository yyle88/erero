package erero

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Error(t, runExample())
}

func runExample() error {
	return op.New("abc")
}
