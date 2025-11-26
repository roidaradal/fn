package ds

type Set[T comparable] struct {
	items map[T]bool
}

// Create new empty set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// Create set from given items
func SetFrom[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	set.AddItems(items)
	return set
}

// Add item to set
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Add items to set
func (s *Set[T]) AddItems(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

// Delete item from set
func (s *Set[T]) Delete(item T) {
	delete(s.items, item)
}

// Check if set contains item
func (s Set[T]) Has(item T) bool {
	return s.items[item]
}

// Check if set doesn't contain item
func (s Set[T]) HasNo(item T) bool {
	return !s.Has(item)
}

// Number of unique set items
func (s Set[T]) Len() int {
	return len(s.items)
}

// Check if set is empty
func (s Set[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Check if set is not empty
func (s Set[T]) NotEmpty() bool {
	return len(s.items) > 0
}

// Return list of set items (random order)
func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

// Compute union of two sets
func (s1 Set[T]) Union(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for _, items := range []map[T]bool{s1.items, s2.items} {
		for item := range items {
			s3.Add(item)
		}
	}
	return s3
}

// Compute intersection of two sets
func (s1 Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for item := range s1.items {
		if s2.Has(item) {
			s3.Add(item)
		}
	}
	return s3
}

// Compute difference of two sets
func (s1 Set[T]) Difference(s2 *Set[T]) *Set[T] {
	s3 := NewSet[T]()
	for item := range s1.items {
		if s2.HasNo(item) {
			s3.Add(item)
		}
	}
	return s3
}
