package ds

import (
	"github.com/roidaradal/fn/dict"
)

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

func SetFrom[T comparable](items []T) *Set[T] {
	s := NewSet[T]()
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) Delete(item T) {
	delete(s.items, item)
}

func (s Set[T]) Contains(item T) bool {
	return dict.HasKey(s.items, item)
}

func (s Set[T]) Len() int {
	return len(s.items)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Set[T]) Items() []T {
	return dict.Keys(s.items)
}

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

func (s1 Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for x := range s2.items {
		if s1.Contains(x) {
			s3.Add(x)
		}
	}
	return s3
}
