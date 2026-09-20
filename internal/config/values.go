package config

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrMissing = errors.New("required config key is missing")

type Values map[string]string

func normalize(key string) string {
	return strings.ReplaceAll(strings.ToLower(key), "-", "_")
}

func (v Values) lookup(key string) (string, bool) {
	raw, ok := v[normalize(key)]
	if !ok || raw == "" {
		return "", false
	}

	return raw, true
}

type parser[T any] func(raw string) (T, error)

func (v Values) Bool(key string, def bool) (bool, error) {
	return optional(v, key, def, strconv.ParseBool)
}

func (v Values) RequireBool(key string) (bool, error) {
	return required(v, key, strconv.ParseBool)
}

func (v Values) Int(key string, def int) (int, error) {
	return optional(v, key, def, strconv.Atoi)
}

func (v Values) RequireInt(key string) (int, error) {
	return required(v, key, strconv.Atoi)
}

func (v Values) String(key string, def string) (string, error) {
	return optional(v, key, def, parseString)
}

func (v Values) RequireString(key string) (string, error) {
	return required(v, key, parseString)
}

func optional[T any](v Values, key string, def T, parse parser[T]) (T, error) {
	raw, ok := v.lookup(key)
	if !ok {
		return def, nil
	}

	return parseValue(key, raw, parse)
}

func required[T any](v Values, key string, parse parser[T]) (T, error) {
	raw, ok := v.lookup(key)
	if !ok {
		var zero T

		return zero, fmt.Errorf("%w: %s", ErrMissing, key)
	}

	return parseValue(key, raw, parse)
}

func parseValue[T any](key string, raw string, parse parser[T]) (T, error) {
	parsed, err := parse(raw)
	if err != nil {
		var zero T

		return zero, fmt.Errorf("config %s: %w", key, err)
	}

	return parsed, nil
}

func parseString(raw string) (string, error) {
	return raw, nil
}
