package flag

import (
	"strings"

	"github.com/knadh/koanf/providers/posflag"
	"github.com/spf13/pflag"

	"github.com/expram/orchestra/internal/config"
)

func Name(path string) string {
	return strings.ReplaceAll(path, config.Delim, "-")
}

func Path(name string) string {
	return strings.ReplaceAll(name, "-", config.Delim)
}

type Flag struct {
	flags *pflag.FlagSet
}

var _ config.Source = (*Flag)(nil)

func New(flags *pflag.FlagSet) *Flag {
	return &Flag{flags: flags}
}

func (s *Flag) Name() string {
	return "flags"
}

func (s *Flag) Read() (map[string]any, error) {
	return posflag.ProviderWithFlag(s.flags, config.Delim, nil, func(flag *pflag.Flag) (string, any) {
		return Path(flag.Name), posflag.FlagVal(s.flags, flag)
	}).Read()
}
