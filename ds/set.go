package ds

import "github.com/roidaradal/fn/dict"

type Set[T comparable] struct {
	items map[T]bool
}

// Creates a new empty set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// Creates a new set from given items
func SetFrom[T comparable](items []T) *Set[T] {
	s := NewSet[T]()
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add item to set
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Delete item from set
func (s *Set[T]) Delete(item T) {
	delete(s.items, item)
}

// Check if set contains item
func (s Set[T]) Contains(item T) bool {
	return dict.HasKey(s.items, item)
}

// Number of unique set items
func (s Set[T]) Len() int {
	return len(s.items)
}

// Check if set is empty
func (s Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Return list of set items
func (s Set[T]) Items() []T {
	return dict.Keys(s.items)
}

// Compute union of two sets
func (s1 Set[T]) Union(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		s3.Add(x)
	}
	for x := range s2.items {
		s3.Add(x)
	}
	return s3
}

// Compute intersection of two sets
func (s1 Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		if s2.Contains(x) {
			s3.Add(x)
		}
	}
	return s3
}

// Compute difference of two sets
func (s1 Set[T]) Difference(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s1.items {
		if !s2.Contains(x) {
			s3.Add(x)
		}
	}
	return s3
}
