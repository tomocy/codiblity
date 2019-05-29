package complexity

import "sort"

func countFrogJumps(x, y, d int) int {
	diff := y - x
	cnt := diff / d
	if diff%d != 0 {
		cnt++
	}

	return cnt
}

func findMissingElement(as []int) int {
	if len(as) == 0 {
		return 1
	}

	sorted := sortByAsc(as)
	var prev, missed int
	for _, a := range sorted {
		if prev+1 != a {
			missed = prev + 1
			break
		}

		prev = a
	}
	if missed == 0 {
		missed = as[len(as)-1] + 1
	}

	return missed
}

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Sort(sort.IntSlice(sorted))

	return sorted
}
