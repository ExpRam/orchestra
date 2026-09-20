package source

import (
	"github.com/expram/orchestra/internal/config"
	"github.com/spf13/pflag"
)

type Flag struct {
	flags *pflag.FlagSet
}

var _ config.Source = (*Flag)(nil)

func NewFlag(flags *pflag.FlagSet) *Flag {
	return &Flag{flags: flags}
}

func (s *Flag) Load() (config.Values, error) {
	values := make(config.Values)

	s.flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Changed {
			values[flag.Name] = flag.Value.String()
		}
	})

	return values, nil
}
