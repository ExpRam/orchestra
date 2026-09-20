package source

import "github.com/expram/orchestra/internal/config"

type Static struct {
	values config.Values
}

var _ config.Source = (*Static)(nil)

func NewStatic(values config.Values) *Static {
	return &Static{values: values}
}

func (s *Static) Load() (config.Values, error) {
	return s.values, nil
}
