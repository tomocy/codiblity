package complexity

import (
	"math"
	"sort"
)

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

func findMinimalDifferenceOfSplitedTapes(as []int) int {
	original := sum(as)
	var min int
	for i := range as {
		if i == 0 {
			continue
		}

		left := as[:i]
		diff := int(math.Abs(float64(original - 2*sum(left))))
		if i == 1 {
			min = diff
			continue
		}
		if diff < min {
			min = diff
		}
	}

	return min
}

func sum(as []int) int {
	var sum int
	for _, a := range as {
		sum += a
	}

	return sum
}
