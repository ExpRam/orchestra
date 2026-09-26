package registry

import (
	"errors"
	"fmt"
)

var ErrAlreadyRegistered = errors.New("already registered")

type Registration[K comparable, V any] struct {
	key   K
	value V
}

func Register[K comparable, V any](key K, value V) Registration[K, V] {
	return Registration[K, V]{key: key, value: value}
}

type Registry[K comparable, V any] struct {
	entries map[K]V
}

func NewRegistry[K comparable, V any](registrations ...Registration[K, V]) *Registry[K, V] {
	entries := make(map[K]V)
	for _, registration := range registrations {
		if err := register(entries, registration); err != nil {
			panic(fmt.Sprintf("%v is registered twice", registration.key))
		}
	}
	return &Registry[K, V]{entries: entries}
}

func NewEmptyRegistry[K comparable, V any]() *Registry[K, V] {
	entries := make(map[K]V)
	return &Registry[K, V]{entries: entries}
}

func register[K comparable, V any](entries map[K]V, registration Registration[K, V]) error {
	if _, ok := entries[registration.key]; ok {
		return ErrAlreadyRegistered
	}

	entries[registration.key] = registration.value

	return nil
}

func (r *Registry[K, V]) Add(registration Registration[K, V]) error {
	return register(r.entries, registration)
}

func (r *Registry[K, V]) Lookup(key K) (V, bool) {
	value, ok := r.entries[key]

	return value, ok
}
