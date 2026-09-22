package reader

import "github.com/expram/orchestra/internal/manifest"

const fileSeparator = ":"

type source struct {
	file  string
	inner manifest.Source
}

var _ manifest.Source = source{}

func newSource(file string, inner manifest.Source) source {
	return source{file: file, inner: inner}
}

func (s source) String() string {
	if s.inner == nil {
		return s.file
	}

	return s.file + fileSeparator + s.inner.String()
}

type Error struct {
	File string
	Err  error
}

func (e Error) Error() string {
	return e.File + fileSeparator + e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
