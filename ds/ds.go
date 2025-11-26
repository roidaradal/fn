// Package ds contains useful data structures.
package ds

type Coords [2]int

// Unpack the coords values
func (c Coords) Tuple() (int, int) {
	return c[0], c[1]
}
