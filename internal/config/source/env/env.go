package env

import (
	"strings"

	"github.com/expram/orchestra/internal/config"
)

type env struct {
	prefix string
}

func (e env) path(name string) string {
	trimmed, ok := strings.CutPrefix(name, e.prefix)
	if !ok || trimmed == "" {
		return ""
	}

	return strings.ReplaceAll(strings.ToLower(trimmed), "_", config.Delim)
}
