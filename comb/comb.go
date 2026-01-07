// Package comb contains combinatorics functions
package comb

import (
	"iter"

	"github.com/roidaradal/fn/list"
	"gonum.org/v1/gonum/stat/combin"
)

// Generator for N-permutations of items]
func Permutations[T any](items []T, take int) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		gen := combin.NewPermutationGenerator(len(items), take)
		i := 0
		for gen.Next() {
			indexes := gen.Permutation(nil)
			combo := createCombo(items, indexes)
			if !yield(i, combo) {
				return
			}
			i += 1
		}
	}
}

// Generator for N-combinations of items
func Combinations[T any](items []T, take int) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		gen := combin.NewCombinationGenerator(len(items), take)
		i := 0
		for gen.Next() {
			indexes := gen.Combination(nil)
			combo := createCombo(items, indexes)
			if !yield(i, combo) {
				return
			}
			i += 1
		}
	}
}

// Generator for all size Permutation of items
func AllPermutations[T any](items []T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		i := 0
		for take := range len(items) + 1 {
			for _, combo := range Permutations(items, take) {
				if !yield(i, combo) {
					return
				}
				i += 1
			}
		}
	}
}

// Generator for all size Combination of items
func AllCombinations[T any](items []T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		i := 0
		for take := range len(items) + 1 {
			for _, combo := range Combinations(items, take) {
				if !yield(i, combo) {
					return
				}
				i += 1
			}
		}
	}
}

// Generator for all size Permutation Positions of items (-1 if not in permutation)
func AllPermutationPositions[T any](items []T) iter.Seq2[int, []int] {
	return func(yield func(int, []int) bool) {
		numItems := len(items)
		indexes := list.NumRange(0, numItems)
		for i, combo := range AllPermutations(indexes) {
			pos := list.Repeated(-1, numItems)
			for p, value := range combo {
				pos[value] = p
			}
			if !yield(i, pos) {
				return
			}
		}
	}
}

// Count permutations(numItems, take)
func NumPermutations(numItems, take int) int {
	return combin.NumPermutations(numItems, take)
}

// Count combinations(numItems, take)
func NumCombinations(numItems, take int) int {
	numerator := Factorial(numItems)
	denom1 := Factorial(take)
	denom2 := Factorial(numItems - take)
	return numerator / (denom1 * denom2)
}

// Common: create combination from items and indexes
func createCombo[T any](items []T, indexes []int) []T {
	combo := make([]T, len(indexes))
	for i, idx := range indexes {
		combo[i] = items[idx]
	}
	return combo
}
