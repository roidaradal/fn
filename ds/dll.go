package ds

// Doubly-Linked list node
type DLLNode[T any] struct {
	Value T
	Next  *DLLNode[T]
	Prev  *DLLNode[T]
}

// Create new DLL node that contains value,
// Prev and Next pointers initialized to self
func NewDLLNode[T any](value T) *DLLNode[T] {
	node := &DLLNode[T]{Value: value}
	node.Next = node
	node.Prev = node
	return node
}

// Remove the DLL node, link its neighbors to fill gap
func (n *DLLNode[T]) Remove() {
	prev, next := n.Prev, n.Next
	prev.Next = next
	next.Prev = prev
	n.Prev = nil
	n.Next = nil
}

// Add new node which contains value before current node
func (n *DLLNode[T]) AddBefore(value T) *DLLNode[T] {
	return addDLLNodeBetween(n.Prev, n, value)
}

// Add new node which contains value after current node
func (n *DLLNode[T]) AddAfter(value T) *DLLNode[T] {
	return addDLLNodeBetween(n, n.Next, value)
}

// Common: add new node which contains value, in between the two nodes
func addDLLNodeBetween[T any](node1, node2 *DLLNode[T], value T) *DLLNode[T] {
	node3 := NewDLLNode(value)
	node1.Next = node3
	node3.Prev = node1
	node2.Prev = node3
	node3.Next = node2
	return node3
}
