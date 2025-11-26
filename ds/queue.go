package ds

import "errors"

var errQueueEmpty = errors.New("queue is empty")

type Queue[T any] struct {
	items []T
}

// Create new empty queue
func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{items: make([]T, 0, capacity)}
}

// Creates queue from given list of items
func QueueFrom[T any](items []T) *Queue[T] {
	return &Queue[T]{items: items}
}

// Add item to end of queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Remove item from front of queue,
// error if empty queue
func (q *Queue[T]) Dequeue() (T, error) {
	front, err := q.Front()
	if err != nil {
		return front, err
	}
	q.items = q.items[1:]
	return front, nil
}

// Return item at front of queue without removing,
// error if empty queue
func (q Queue[T]) Front() (T, error) {
	var front T
	if q.IsEmpty() {
		return front, errQueueEmpty
	}
	front = q.items[0]
	return front, nil
}

// Number of queue items left
func (q Queue[T]) Len() int {
	return len(q.items)
}

// Checks if queue is empty
func (q Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Checks if queue is not empty
func (q Queue[T]) NotEmpty() bool {
	return len(q.items) > 0
}

// Return list of queue items
func (q Queue[T]) Items() []T {
	return q.items
}
