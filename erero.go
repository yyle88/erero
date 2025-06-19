package erero

func Ero(err error) error {
	return ero.Ero(err)
}

func New(message string) error {
	return ero.New(message)
}

func WithMessage(err error, message string) error {
	return ero.WithMessage(err, message)
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return ero.WithMessagef(err, format, args...)
}

func Errorf(format string, args ...interface{}) error {
	return ero.Errorf(format, args...)
}

func Is(err, target error) bool {
	return ero.Is(err, target)
}

func Ise(err, target error) bool {
	return ero.Ise(err, target)
}

func As(err error, target any) bool {
	return ero.As(err, target)
}

func Ase(err error, target any) bool {
	return ero.Ase(err, target)
}

func Wro(err error) error {
	return ero.Wro(err)
}

func Wrap(err error, message string) error {
	return ero.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return ero.Wrapf(err, format, args...)
}

func Join(errs ...error) error {
	return ero.Join(errs...)
}

func Joins(errs []error) error {
	return ero.Joins(errs)
}
