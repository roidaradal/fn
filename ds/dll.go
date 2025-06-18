package ds

type DLLNode[T any] struct {
	Value T
	Next  *DLLNode[T]
	Prev  *DLLNode[T]
}

func NewDLLNode[T any](value T) *DLLNode[T] {
	node := &DLLNode[T]{Value: value}
	node.Next = node
	node.Prev = node
	return node
}

func (n1 *DLLNode[T]) AddBetween(n2 *DLLNode[T], value T) *DLLNode[T] {
	n3 := NewDLLNode(value)
	n1.Next = n3
	n3.Prev = n1
	n2.Prev = n3
	n3.Next = n2
	return n3
}

func (n *DLLNode[T]) Remove() {
	prev := n.Prev
	next := n.Next
	prev.Next = next
	next.Prev = prev
	n.Prev = nil
	n.Next = nil
}
