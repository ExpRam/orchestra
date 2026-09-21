package source

import (
	"fmt"

	"github.com/expram/orchestra/internal/config"
	"github.com/spf13/pflag"
)

type Flag struct {
	flags *pflag.FlagSet
	args  []string
}

var _ config.Source = (*Flag)(nil)

func NewFlag(flags *pflag.FlagSet, args []string) *Flag {
	return &Flag{flags: flags, args: args}
}

func (s *Flag) Load() (config.Values, error) {
	unknown := s.flags.ParseErrorsAllowlist.UnknownFlags
	s.flags.ParseErrorsAllowlist.UnknownFlags = true
	defer func() { s.flags.ParseErrorsAllowlist.UnknownFlags = unknown }()

	if err := s.flags.Parse(s.args); err != nil {
		return nil, fmt.Errorf("parse flags: %w", err)
	}

	values := make(config.Values)

	s.flags.Visit(func(flag *pflag.Flag) {
		values[flag.Name] = flag.Value.String()
	})

	return values, nil
}
