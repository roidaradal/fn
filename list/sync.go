package list

import "sync"

// Concurrency-safe generic list
type SyncList[T any] struct {
	mu    sync.RWMutex
	items []T
}

// Create new SyncList
func NewSyncList[T any]() *SyncList[T] {
	return &SyncList[T]{items: make([]T, 0)}
}

// Create SyncList from existing list
func SyncListFrom[T any](items []T) *SyncList[T] {
	return &SyncList[T]{items: items}
}

// SyncList.Append
func (sl *SyncList[T]) Append(items ...T) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.items = append(sl.items, items...)
}

// SyncList.Clear
func (sl *SyncList[T]) Clear() {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	sl.items = make([]T, 0)
}

// SyncList number of items
func (sl *SyncList[T]) Len() int {
	sl.mu.RLock()
	defer sl.mu.RUnlock()
	return len(sl.items)
}

// SyncList underlying items
func (sl *SyncList[T]) Items() []T {
	sl.mu.RLock()
	defer sl.mu.RUnlock()
	return Copy(sl.items)
}

// Copy SyncList's items and clear it
func (sl *SyncList[T]) ClearItems() []T {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	items := Copy(sl.items)
	sl.items = make([]T, 0)
	return items
}
