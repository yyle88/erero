package errorb

import "github.com/pkg/errors"

type ErrorOrNotExist struct {
	ErrCause error
	NotExist bool
}

func NewErrorOrNotExist(errCause error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		ErrCause: errCause,
		NotExist: false,
	}
}

func NewErrorIsNotExist() *ErrorOrNotExist {
	return &ErrorOrNotExist{
		ErrCause: errors.New("not-exist"),
		NotExist: true,
	}
}

type ErrorOrSuccess struct {
	ErrCause error
	Success  bool
}

func NewErrorOrSuccess(errCause error) *ErrorOrSuccess {
	return &ErrorOrSuccess{
		ErrCause: errCause,
		Success:  false,
	}
}

func NewErrorIsSuccess() *ErrorOrSuccess {
	return &ErrorOrSuccess{
		ErrCause: errors.New("success"),
		Success:  true,
	}
}

type ErrorOrIgnore struct {
	ErrCause error
	Ignore   bool
}

func NewErrorOrIgnore(errCause error) *ErrorOrIgnore {
	return &ErrorOrIgnore{
		ErrCause: errCause,
		Ignore:   false,
	}
}

func NewErrorIsIgnore() *ErrorOrIgnore {
	return &ErrorOrIgnore{
		ErrCause: errors.New("ignore"),
		Ignore:   true,
	}
}

type ErrorOrRetryable struct {
	ErrCause  error
	Retryable bool
}

func NewErrorOrRetryable(errCause error) *ErrorOrRetryable {
	return &ErrorOrRetryable{
		ErrCause:  errCause,
		Retryable: false,
	}
}

func NewErrorIsRetryable() *ErrorOrRetryable {
	return &ErrorOrRetryable{
		ErrCause:  errors.New("retryable"),
		Retryable: true,
	}
}

type ErrorOrRecoverable struct {
	ErrCause    error
	Recoverable bool
}

func NewErrorOrRecoverable(errCause error) *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		ErrCause:    errCause,
		Recoverable: false,
	}
}

func NewErrorIsRecoverable() *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		ErrCause:    errors.New("recoverable"),
		Recoverable: true,
	}
}
