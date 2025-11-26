package dict

import "sync"

// Concurrency-safe generic map
type SyncMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// Create new SyncMap
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{data: make(map[K]V)}
}

// Create SyncMap from existing map
func SyncMapFrom[K comparable, V any](items map[K]V) *SyncMap[K, V] {
	return &SyncMap[K, V]{data: items}
}

// SyncMap.Get
func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	return value, ok
}

// SyncMap.Set
func (sm *SyncMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// SyncMap.Delete
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Clear SyncMap's underlying map data
func (sm *SyncMap[K, V]) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	clear(sm.data)
}

// SyncMap number of items
func (sm *SyncMap[K, V]) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// SyncMap's underlying map
func (sm *SyncMap[K, V]) Map() map[K]V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data
}

// SyncMap's underlying map keys
func (sm *SyncMap[K, V]) Keys() []K {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return Keys(sm.data)
}

// SyncMap's underlying map values
func (sm *SyncMap[K, V]) Values() []V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return Values(sm.data)
}
