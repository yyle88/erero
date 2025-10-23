package errlog_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoin(t *testing.T) {
	require.Error(t, caseEro.Join(
		errors.New("abc"),
		errors.New("xyz"),
		errors.New("123"),
	))
}

func TestJoins(t *testing.T) {
	var errs = []error{
		errors.New("abc"),
		errors.New("xyz"),
		errors.New("123"),
	}
	require.Error(t, caseEro.Joins(errs))
}
