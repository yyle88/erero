package erero

func Ero(erx error) error {
	return op.Ero(erx)
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

func Is(erx, target error) bool {
	return op.Is(erx, target)
}

func Ise(erx, target error) bool {
	return op.Ise(erx, target)
}

func As(erx error, target any) bool {
	return op.As(erx, target)
}

func Ase(erx error, target any) bool {
	return op.Ase(erx, target)
}

func Wro(erx error) error {
	return op.Wro(erx)
}

func Wrap(err error, message string) error {
	return op.Wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return op.Wrapf(err, format, args...)
}
