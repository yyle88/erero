package errorb

import "github.com/pkg/errors"

type ErrorOrNotExist struct {
	Err      error
	NotExist bool
}

func NewErrorOrNotExist(err error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Err:      err,
		NotExist: false,
	}
}

func NewErrorIsNotExist() *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Err:      errors.New("not-exist"),
		NotExist: true,
	}
}

type ErrorOrSuccess struct {
	Err     error
	Success bool
}

func NewErrorOrSuccess(err error) *ErrorOrSuccess {
	return &ErrorOrSuccess{
		Err:     err,
		Success: false,
	}
}

func NewErrorIsSuccess() *ErrorOrSuccess {
	return &ErrorOrSuccess{
		Err:     errors.New("success"),
		Success: true,
	}
}

type ErrorOrIgnore struct {
	Err    error
	Ignore bool
}

func NewErrorOrIgnore(err error) *ErrorOrIgnore {
	return &ErrorOrIgnore{
		Err:    err,
		Ignore: false,
	}
}

func NewErrorIsIgnore() *ErrorOrIgnore {
	return &ErrorOrIgnore{
		Err:    errors.New("ignore"),
		Ignore: true,
	}
}

type ErrorOrRetryable struct {
	Err       error
	Retryable bool
}

func NewErrorOrRetryable(err error) *ErrorOrRetryable {
	return &ErrorOrRetryable{
		Err:       err,
		Retryable: false,
	}
}

func NewErrorIsRetryable() *ErrorOrRetryable {
	return &ErrorOrRetryable{
		Err:       errors.New("retryable"),
		Retryable: true,
	}
}

type ErrorOrRecoverable struct {
	Err         error
	Recoverable bool
}

func NewErrorOrRecoverable(err error) *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		Err:         err,
		Recoverable: false,
	}
}

func NewErrorIsRecoverable() *ErrorOrRecoverable {
	return &ErrorOrRecoverable{
		Err:         errors.New("recoverable"),
		Recoverable: true,
	}
}
