package config

type Source interface {
	Name() string
	Read() (map[string]any, error)
}
