package dict

import "sync"

// Concurrency-safe generic map
type SyncMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// Creates a new SyncMap
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{data: make(map[K]V)}
}

// Creates a SyncMap from existing map
func SyncMapFrom[K comparable, V any](items map[K]V) *SyncMap[K, V] {
	return &SyncMap[K, V]{data: items}
}

// Concurrent-safe map setter
func (sm *SyncMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Concurrent-safe map getter
func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	return value, ok
}

// Concurrent-safe map delete key
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Clear underlying map's data
func (sm *SyncMap[K, V]) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	clear(sm.data)
}

// Get underlying map from SyncMap
func (sm *SyncMap[K, V]) Map() map[K]V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data
}

// Get underlying map's keys
func (sm *SyncMap[K, V]) Keys() []K {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return Keys(sm.data)
}

// Get underlying map's values
func (sm *SyncMap[K, V]) Values() []V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return Values(sm.data)
}
