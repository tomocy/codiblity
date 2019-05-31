package sort

import "sort"

func countDistincts(as []int) int {
	dists := make(map[int]int)
	for _, a := range as {
		dists[a]++
	}

	return len(dists)
}

func findMaximalProduct(as []int) int {
	sorted := sortByAsc(as)
	smallest, smallest2 := sorted[0], sorted[1]
	largest, largest2, largest3 := sorted[len(sorted)-1], sorted[len(sorted)-2], sorted[len(sorted)-3]
	if 0 <= largest && largest2*largest3 < smallest*smallest2 {
		return smallest * smallest2 * largest
	}

	return largest3 * largest2 * largest
}

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Ints(sorted)

	return sorted
}
