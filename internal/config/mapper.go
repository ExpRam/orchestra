package config

func ToConfig(v Values) (Config, error) {
	debug, err := v.Bool(KeyDebug, false)
	if err != nil {
		return Config{}, err
	}

	return Config{Debug: debug}, nil
}
