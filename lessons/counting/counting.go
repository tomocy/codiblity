package counting

import "sort"

func isPermutation(as []int) bool {
	sorted := sortByAsc(as)
	var prev int
	for _, a := range sorted {
		if prev+1 != a {
			return false
		}

		prev = a
	}

	return true
}

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Sort(sort.IntSlice(sorted))

	return sorted
}
