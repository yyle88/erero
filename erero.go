package erero

func Ero(err error) error {
	return op.Ero(err)
}

func New(message string) error {
	return op.New(message)
}

func WithMessage(err error, message string) error {
	return op.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return op.WithMessagef(err, format, args...)
}

func Errorf(format string, args ...interface{}) error {
	return op.Errorf(format, args...)
}

func Is(err, target error) bool {
	return op.Is(err, target)
}

func Ise(err, target error) bool {
	return op.Ise(err, target)
}

func As(err error, target any) bool {
	return op.As(err, target)
}

func Ase(err error, target any) bool {
	return op.Ase(err, target)
}

func Wro(err error) error {
	return op.Wro(err)
}

func Wrap(err error, message string) error {
	return op.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return op.Wrapf(err, format, args...)
}
