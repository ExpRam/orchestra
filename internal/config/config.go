package config

const Delim = "."

const PathDebug = "debug"

type Config struct {
	Debug bool `config:"debug"`
}

func Defaults() Config {
	return Config{Debug: false}
}
