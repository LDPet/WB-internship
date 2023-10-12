package amcacher

import (
	"golang.org/x/exp/maps"
	"sync"
)

type AsyncMapCacher[K comparable, V any] struct {
	hashMap map[K]V
	mutex   sync.RWMutex
}

func NewAsyncMapCacher[K comparable, V any]() *AsyncMapCacher[K, V] {
	return &AsyncMapCacher[K, V]{
		hashMap: make(map[K]V),
	}
}

func (a *AsyncMapCacher[K, V]) Set(key K, value V) {
	a.mutex.Lock()
	a.hashMap[key] = value
	a.mutex.Unlock()
}

func (a *AsyncMapCacher[K, V]) Get(key K) (V, bool) {
	a.mutex.RLock()
	val, ok := a.hashMap[key]
	a.mutex.RUnlock()

	return val, ok
}

func (a *AsyncMapCacher[K, V]) Delete(key K) {
	a.mutex.Lock()
	delete(a.hashMap, key)
	a.mutex.Unlock()
}

func (a *AsyncMapCacher[K, V]) Keys() []K {
	a.mutex.RLock()
	keys := maps.Keys(a.hashMap)
	a.mutex.RUnlock()
	return keys
}
