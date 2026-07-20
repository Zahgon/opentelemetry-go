package metric

import (
	"sync"
)

type cache[K comparable, V any] struct {
	sync.Mutex
	data map[K]V
}

func (c *cache[K, V]) Lookup(key K, f func() V) V { _ = "STUB: not implemented"; return *new(V) }

func (c *cache[K, V]) HasKey(key K) bool { _ = "STUB: not implemented"; return false }

type cacheWithErr[K comparable, V any] struct {
	cache[K, valAndErr[V]]
}

type valAndErr[V any] struct {
	val V
	err error
}

func (c *cacheWithErr[K, V]) Lookup(key K, f func() (V, error)) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}
