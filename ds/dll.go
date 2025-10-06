package ds

// Doubly-linked list node
type DLLNode[T any] struct {
	Value T
	Next  *DLLNode[T]
	Prev  *DLLNode[T]
}

// Create new DLL node that contains given value,
// Prev and next pointers are initialized to self
func NewDLLNode[T any](value T) *DLLNode[T] {
	node := &DLLNode[T]{Value: value}
	node.Next = node
	node.Prev = node
	return node
}

// Add a new DLL node in between node1 and node2, which contains the given value,
// Assumed that node1 <=> node2 are linked to each other
func (n1 *DLLNode[T]) AddBetween(n2 *DLLNode[T], value T) *DLLNode[T] {
	n3 := NewDLLNode(value)
	n1.Next = n3
	n3.Prev = n1
	n2.Prev = n3
	n3.Next = n2
	return n3
}

// Removes the DLL node, linking its neighbors to fill the gap
func (n *DLLNode[T]) Remove() {
	prev, next := n.Prev, n.Next
	prev.Next = next
	next.Prev = prev
	n.Prev = nil
	n.Next = nil
}
