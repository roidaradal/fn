package ds

import "errors"

var errQueueEmpty = errors.New("queue is empty")

type Queue[T any] struct {
	items []T
}

// Creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Creates a queue from the given items
func QueueFrom[T any](items []T) *Queue[T] {
	q := NewQueue[T]()
	q.items = items
	return q
}

// Add item to end of the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Remove item from the front of the queue,
// Error if queue is empty
func (q *Queue[T]) Dequeue() (T, error) {
	front, err := q.Front()
	if err != nil {
		return front, err
	}
	q.items = q.items[1:]
	return front, nil
}

// Return the item at front of queue without removing it,
// Error if queue is empty
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

// Return list of queue items
func (q Queue[T]) Items() []T {
	return q.items
}
