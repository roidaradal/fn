package dict

import "sync"

// Concurrency-safe generic map
type SyncMap[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

// Creates a new SyncMap
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{data: make(map[K]V)}
}

// Concurrent-safe map setter
func (sm *SyncMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Concurrent-safe map getter
func (sm *SyncMap[K, V]) Get(key K) (V, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.data[key]
	return value, ok
}

// Concurrent-safe map delete key
func (sm *SyncMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}
