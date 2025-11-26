package ds

import "strings"

// Struct for list of items
type List[T any] struct {
	Items []T
	Count int
}

// Create new List from list of items
func NewList[T any](items []T) *List[T] {
	if items == nil {
		items = []T{}
	}
	return &List[T]{Items: items, Count: len(items)}
}

// Struct for list of items, with associated lookup table,
// T = item type, L = lookup item type
type ListWithLookup[T, L any] struct {
	Items  []T
	Lookup map[string]L
	Count  int
}

// Create new ListWithLookup from given list of items,
// Initialize lookup to empty map
func NewListWithLookup[T, L any](items []T) *ListWithLookup[T, L] {
	if items == nil {
		items = []T{}
	}
	return &ListWithLookup[T, L]{
		Items:  items,
		Lookup: make(map[string]L),
		Count:  len(items),
	}
}

// Struct for map of items, with associated lookup table,
// K, V = map key/value types, L = lookup item type
type MapWithLookup[K comparable, V, L any] struct {
	Items  map[K]V
	Lookup map[string]L
	Count  int
}

// Create new MapWithLookup from given map of items,
// Initialize lookup to empty map
func NewMapWithLookup[K comparable, V, L any](items map[K]V) *MapWithLookup[K, V, L] {
	if items == nil {
		items = make(map[K]V)
	}
	return &MapWithLookup[K, V, L]{
		Items:  items,
		Lookup: make(map[string]L),
		Count:  len(items),
	}
}

// Struct for lines of data rows, separated by comma
type DataRows struct {
	Rows string
}

// Get lines from rows, and split each line by comma
func (d DataRows) GetRows() [][]string {
	lines := strings.Split(d.Rows, "\n")
	rows := make([][]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		row := strings.Split(line, ",")
		for i, cell := range row {
			row[i] = strings.TrimSpace(cell)
		}
		rows = append(rows, row)
	}
	return rows
}
