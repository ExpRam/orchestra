package errscope

const scopeSeparator = ": "

type Error struct {
	Scope string
	Cause error
}

func In(scope string, cause error) error {
	return Error{Scope: scope, Cause: cause}
}

func (e Error) Error() string {
	return e.Scope + scopeSeparator + e.Cause.Error()
}

func (e Error) Unwrap() error {
	return e.Cause
}
