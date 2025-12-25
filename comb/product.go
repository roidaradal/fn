package comb

import (
	"iter"

	"github.com/roidaradal/fn/list"
)

// Create an iterator for the Cartesian product of given domains
func Product[T any](domains ...[]T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		sizes := list.Map(domains, list.Length)
		total := list.Product(sizes)
		for i := range total {
			tuple := productDomainCombo(domains, sizes, i)
			if !yield(i, tuple) {
				return
			}
		}
	}
}

// Create an iterator for Cartesian product of given domains, from index [start, end)
func RangeProduct[T any](start, end int, domains ...[]T) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		sizes := list.Map(domains, list.Length)
		for i := start; i < end; i++ {
			tuple := productDomainCombo(domains, sizes, i)
			if !yield(i, tuple) {
				return
			}
		}
	}
}

// Utility: computes the Cartesian product combination from given domains for given index
func productDomainCombo[T any](domains [][]T, sizes []int, index int) []T {
	numSizes := len(sizes)
	indexes := make([]int, numSizes)
	for i := range numSizes {
		denom := list.Product(sizes[i+1:])
		num := sizes[i] * denom
		indexes[i] = (index % num) / denom
	}
	numDomains := len(domains)
	combo := make([]T, numDomains)
	for i := range numDomains {
		combo[i] = domains[i][indexes[i]]
	}
	return combo
}
