package ds

import (
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/str"
)

// Object wrapper for list of items
type List[T any] struct {
	Items []T
	Count int
}

// Create new List object from given list of items
func NewList[T any](items []T) *List[T] {
	if items == nil {
		items = []T{}
	}
	return &List[T]{Items: items, Count: len(items)}
}

// List of items with an associated lookup table,
// T = item type, L = lookup item type
type ListLookup[T any, L any] struct {
	Items  []T
	Lookup map[string]L
	Count  int
}

// Create new ListLookup object from given list of items,
// Initialize lookup to empty map
func NewListLookup[T any, L any](items []T) *ListLookup[T, L] {
	if items == nil {
		items = []T{}
	}
	return &ListLookup[T, L]{
		Items:  items,
		Lookup: make(map[string]L),
		Count:  len(items),
	}
}

// Map of items with an associated lookup table,
// T = value type of items map, L = lookup item type
type MapLookup[T any, L any] struct {
	Items  map[string]T
	Lookup map[string]L
	Count  int
}

// Create new MapLookup object from given map of items,
// Initialize lookup to empty map
func NewMapLookup[T any, L any](items map[string]T) *MapLookup[T, L] {
	if items == nil {
		items = make(map[string]T)
	}
	return &MapLookup[T, L]{
		Items:  items,
		Lookup: make(map[string]L),
		Count:  len(items),
	}
}

// Object wrapper for lines of data rows, separated by comma
type DataRows struct {
	Rows string
}

// Get the lines from Rows, and split each line by comma
func (d DataRows) GetRows() [][]string {
	lines := str.Lines(d.Rows)
	return fn.Map(lines, str.CommaSplit)
}
