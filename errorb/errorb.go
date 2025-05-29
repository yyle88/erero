package errorb

import "github.com/pkg/errors"

type ErrorOrNotExist struct {
	Cause    error
	NotExist bool
}

func NewErrorOrNotExist(cause error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Cause:    cause,
		NotExist: false,
	}
}

func NewErrorIsNotExist() *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Cause:    errors.New("NOT-EXIST"),
		NotExist: true,
	}
}

type ErrorOrSuccess struct {
	Cause   error
	Success bool
}

func NewErrorOrSuccess(cause error) *ErrorOrSuccess {
	return &ErrorOrSuccess{
		Cause:   cause,
		Success: false,
	}
}

func NewErrorIsSuccess() *ErrorOrSuccess {
	return &ErrorOrSuccess{
		Cause:   errors.New("SUCCESS"),
		Success: true,
	}
}

type ErrorOrIgnore struct {
	Cause  error
	Ignore bool
}

func NewErrorOrIgnore(cause error) *ErrorOrIgnore {
	return &ErrorOrIgnore{
		Cause:  cause,
		Ignore: false,
	}
}

func NewErrorIsIgnore() *ErrorOrIgnore {
	return &ErrorOrIgnore{
		Cause:  errors.New("IGNORE"),
		Ignore: true,
	}
}

type ErrorOrRetryable struct {
	Cause     error
	Retryable bool
}

func NewErrorOrRetryable(cause error) *ErrorOrRetryable {
	return &ErrorOrRetryable{
		Cause:     cause,
		Retryable: false,
	}
}

func NewErrorIsRetryable() *ErrorOrRetryable {
	return &ErrorOrRetryable{
		Cause:     errors.New("RETRYABLE"),
		Retryable: true,
	}
}

type ErrorOrRecoverable struct {
	Cause       error
	Recoverable bool
}

func NewErrorOrRecoverable(cause error) *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		Cause:       cause,
		Recoverable: false,
	}
}

func NewErrorIsRecoverable() *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		Cause:       errors.New("RECOVERABLE"),
		Recoverable: true,
	}
}
