package registry

import "fmt"

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

func NewRegistry[K comparable, V any](registrations ...Registration[K, V]) Registry[K, V] {
	entries := make(map[K]V, len(registrations))

	for _, registration := range registrations {
		if _, ok := entries[registration.key]; ok {
			panic(fmt.Sprintf("%v is registered twice", registration.key))
		}

		entries[registration.key] = registration.value
	}

	return Registry[K, V]{entries: entries}
}

func (r Registry[K, V]) Lookup(key K) (V, bool) {
	value, ok := r.entries[key]

	return value, ok
}
