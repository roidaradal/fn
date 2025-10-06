package ds

import "errors"

var errStackEmpty = errors.New("stack is empty")

type Stack[T any] struct {
	items []T
}

// Creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Creates a stack from the given items (last item = top)
func StackFrom[T any](items []T) *Stack[T] {
	s := NewStack[T]()
	s.items = items
	return s
}

// Push item to top of stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Remove item from the top of stack,
// Error if stack is empty
func (s *Stack[T]) Pop() (T, error) {
	top, err := s.Top()
	if err != nil {
		return top, err
	}
	last := len(s.items) - 1
	s.items = s.items[:last]
	return top, nil
}

// Return the item at the top of stack without removing it,
// Error if stack is empty
func (s Stack[T]) Top() (T, error) {
	var top T
	if s.IsEmpty() {
		return top, errStackEmpty
	}
	top = s.items[len(s.items)-1]
	return top, nil
}

// Number of stack items left
func (s Stack[T]) Len() int {
	return len(s.items)
}

// Check if stack is empty
func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Return list of stack items
func (s Stack[T]) Items() []T {
	return s.items
}
