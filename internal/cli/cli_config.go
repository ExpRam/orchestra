package config

import (
	"github.com/spf13/pflag"
)

type Command struct {
	flags *pflag.FlagSet
}

func NewCommand(flags *pflag.FlagSet) *Command {
	return &Command{
		flags: flags,
	}
}

func (s *Command) Load() (Values, error) {
	values := make(Values)

	s.flags.Visit(func(flag *pflag.Flag) {
		values[flag.Name] = flag.Value.String()
	})

	return values, nil
}
